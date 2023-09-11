/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package logging_banzaicloud_io_v1beta1

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
	_ datasource.DataSource = &LoggingBanzaicloudIoSyslogNgflowV1Beta1Manifest{}
)

func NewLoggingBanzaicloudIoSyslogNgflowV1Beta1Manifest() datasource.DataSource {
	return &LoggingBanzaicloudIoSyslogNgflowV1Beta1Manifest{}
}

type LoggingBanzaicloudIoSyslogNgflowV1Beta1Manifest struct{}

type LoggingBanzaicloudIoSyslogNgflowV1Beta1ManifestData struct {
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
		Filters *[]struct {
			Id    *string `tfsdk:"id" json:"id,omitempty"`
			Match *struct {
				And    *map[string]string `tfsdk:"and" json:"and,omitempty"`
				Not    *map[string]string `tfsdk:"not" json:"not,omitempty"`
				Or     *map[string]string `tfsdk:"or" json:"or,omitempty"`
				Regexp *struct {
					Flags    *[]string `tfsdk:"flags" json:"flags,omitempty"`
					Pattern  *string   `tfsdk:"pattern" json:"pattern,omitempty"`
					Template *string   `tfsdk:"template" json:"template,omitempty"`
					Type     *string   `tfsdk:"type" json:"type,omitempty"`
					Value    *string   `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"regexp" json:"regexp,omitempty"`
			} `tfsdk:"match" json:"match,omitempty"`
			Parser *struct {
				Metrics_probe *struct {
					Key    *string            `tfsdk:"key" json:"key,omitempty"`
					Labels *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					Level  *int64             `tfsdk:"level" json:"level,omitempty"`
				} `tfsdk:"metrics_probe" json:"metrics-probe,omitempty"`
				Regexp *struct {
					Flags    *[]string `tfsdk:"flags" json:"flags,omitempty"`
					Patterns *[]string `tfsdk:"patterns" json:"patterns,omitempty"`
					Prefix   *string   `tfsdk:"prefix" json:"prefix,omitempty"`
					Template *string   `tfsdk:"template" json:"template,omitempty"`
				} `tfsdk:"regexp" json:"regexp,omitempty"`
				Syslog_parser *struct {
					Flags *[]string `tfsdk:"flags" json:"flags,omitempty"`
				} `tfsdk:"syslog_parser" json:"syslog-parser,omitempty"`
			} `tfsdk:"parser" json:"parser,omitempty"`
			Rewrite *[]struct {
				Group_unset *struct {
					Condition *struct {
						And    *map[string]string `tfsdk:"and" json:"and,omitempty"`
						Not    *map[string]string `tfsdk:"not" json:"not,omitempty"`
						Or     *map[string]string `tfsdk:"or" json:"or,omitempty"`
						Regexp *struct {
							Flags    *[]string `tfsdk:"flags" json:"flags,omitempty"`
							Pattern  *string   `tfsdk:"pattern" json:"pattern,omitempty"`
							Template *string   `tfsdk:"template" json:"template,omitempty"`
							Type     *string   `tfsdk:"type" json:"type,omitempty"`
							Value    *string   `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"regexp" json:"regexp,omitempty"`
					} `tfsdk:"condition" json:"condition,omitempty"`
					Pattern *string `tfsdk:"pattern" json:"pattern,omitempty"`
				} `tfsdk:"group_unset" json:"group_unset,omitempty"`
				Rename *struct {
					Condition *struct {
						And    *map[string]string `tfsdk:"and" json:"and,omitempty"`
						Not    *map[string]string `tfsdk:"not" json:"not,omitempty"`
						Or     *map[string]string `tfsdk:"or" json:"or,omitempty"`
						Regexp *struct {
							Flags    *[]string `tfsdk:"flags" json:"flags,omitempty"`
							Pattern  *string   `tfsdk:"pattern" json:"pattern,omitempty"`
							Template *string   `tfsdk:"template" json:"template,omitempty"`
							Type     *string   `tfsdk:"type" json:"type,omitempty"`
							Value    *string   `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"regexp" json:"regexp,omitempty"`
					} `tfsdk:"condition" json:"condition,omitempty"`
					NewName *string `tfsdk:"new_name" json:"newName,omitempty"`
					OldName *string `tfsdk:"old_name" json:"oldName,omitempty"`
				} `tfsdk:"rename" json:"rename,omitempty"`
				Set *struct {
					Condition *struct {
						And    *map[string]string `tfsdk:"and" json:"and,omitempty"`
						Not    *map[string]string `tfsdk:"not" json:"not,omitempty"`
						Or     *map[string]string `tfsdk:"or" json:"or,omitempty"`
						Regexp *struct {
							Flags    *[]string `tfsdk:"flags" json:"flags,omitempty"`
							Pattern  *string   `tfsdk:"pattern" json:"pattern,omitempty"`
							Template *string   `tfsdk:"template" json:"template,omitempty"`
							Type     *string   `tfsdk:"type" json:"type,omitempty"`
							Value    *string   `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"regexp" json:"regexp,omitempty"`
					} `tfsdk:"condition" json:"condition,omitempty"`
					Field *string `tfsdk:"field" json:"field,omitempty"`
					Value *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"set" json:"set,omitempty"`
				Subst *struct {
					Condition *struct {
						And    *map[string]string `tfsdk:"and" json:"and,omitempty"`
						Not    *map[string]string `tfsdk:"not" json:"not,omitempty"`
						Or     *map[string]string `tfsdk:"or" json:"or,omitempty"`
						Regexp *struct {
							Flags    *[]string `tfsdk:"flags" json:"flags,omitempty"`
							Pattern  *string   `tfsdk:"pattern" json:"pattern,omitempty"`
							Template *string   `tfsdk:"template" json:"template,omitempty"`
							Type     *string   `tfsdk:"type" json:"type,omitempty"`
							Value    *string   `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"regexp" json:"regexp,omitempty"`
					} `tfsdk:"condition" json:"condition,omitempty"`
					Field   *string   `tfsdk:"field" json:"field,omitempty"`
					Flags   *[]string `tfsdk:"flags" json:"flags,omitempty"`
					Pattern *string   `tfsdk:"pattern" json:"pattern,omitempty"`
					Replace *string   `tfsdk:"replace" json:"replace,omitempty"`
					Type    *string   `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"subst" json:"subst,omitempty"`
				Unset *struct {
					Condition *struct {
						And    *map[string]string `tfsdk:"and" json:"and,omitempty"`
						Not    *map[string]string `tfsdk:"not" json:"not,omitempty"`
						Or     *map[string]string `tfsdk:"or" json:"or,omitempty"`
						Regexp *struct {
							Flags    *[]string `tfsdk:"flags" json:"flags,omitempty"`
							Pattern  *string   `tfsdk:"pattern" json:"pattern,omitempty"`
							Template *string   `tfsdk:"template" json:"template,omitempty"`
							Type     *string   `tfsdk:"type" json:"type,omitempty"`
							Value    *string   `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"regexp" json:"regexp,omitempty"`
					} `tfsdk:"condition" json:"condition,omitempty"`
					Field *string `tfsdk:"field" json:"field,omitempty"`
				} `tfsdk:"unset" json:"unset,omitempty"`
			} `tfsdk:"rewrite" json:"rewrite,omitempty"`
		} `tfsdk:"filters" json:"filters,omitempty"`
		GlobalOutputRefs *[]string `tfsdk:"global_output_refs" json:"globalOutputRefs,omitempty"`
		LocalOutputRefs  *[]string `tfsdk:"local_output_refs" json:"localOutputRefs,omitempty"`
		LoggingRef       *string   `tfsdk:"logging_ref" json:"loggingRef,omitempty"`
		Match            *struct {
			And    *map[string]string `tfsdk:"and" json:"and,omitempty"`
			Not    *map[string]string `tfsdk:"not" json:"not,omitempty"`
			Or     *map[string]string `tfsdk:"or" json:"or,omitempty"`
			Regexp *struct {
				Flags    *[]string `tfsdk:"flags" json:"flags,omitempty"`
				Pattern  *string   `tfsdk:"pattern" json:"pattern,omitempty"`
				Template *string   `tfsdk:"template" json:"template,omitempty"`
				Type     *string   `tfsdk:"type" json:"type,omitempty"`
				Value    *string   `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"regexp" json:"regexp,omitempty"`
		} `tfsdk:"match" json:"match,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *LoggingBanzaicloudIoSyslogNgflowV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_logging_banzaicloud_io_syslog_ng_flow_v1beta1_manifest"
}

func (r *LoggingBanzaicloudIoSyslogNgflowV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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
				Description:         "",
				MarkdownDescription: "",
				Attributes: map[string]schema.Attribute{
					"filters": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"match": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"and": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"not": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"or": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"regexp": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"flags": schema.ListAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"pattern": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"template": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"type": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"value": schema.StringAttribute{
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
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"parser": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"metrics_probe": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"key": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"labels": schema.MapAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"level": schema.Int64Attribute{
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

										"regexp": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"flags": schema.ListAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"patterns": schema.ListAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"prefix": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"template": schema.StringAttribute{
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

										"syslog_parser": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"flags": schema.ListAttribute{
													Description:         "",
													MarkdownDescription: "",
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

								"rewrite": schema.ListNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"group_unset": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"condition": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"and": schema.MapAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"not": schema.MapAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"or": schema.MapAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"regexp": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"flags": schema.ListAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"pattern": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"template": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"type": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"value": schema.StringAttribute{
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pattern": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"rename": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"condition": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"and": schema.MapAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"not": schema.MapAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"or": schema.MapAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"regexp": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"flags": schema.ListAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"pattern": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"template": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"type": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"value": schema.StringAttribute{
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"new_name": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"old_name": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"set": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"condition": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"and": schema.MapAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"not": schema.MapAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"or": schema.MapAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"regexp": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"flags": schema.ListAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"pattern": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"template": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"type": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"value": schema.StringAttribute{
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"field": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"value": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"subst": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"condition": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"and": schema.MapAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"not": schema.MapAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"or": schema.MapAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"regexp": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"flags": schema.ListAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"pattern": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"template": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"type": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"value": schema.StringAttribute{
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"field": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"flags": schema.ListAttribute{
														Description:         "",
														MarkdownDescription: "",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"pattern": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"replace": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"type": schema.StringAttribute{
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

											"unset": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"condition": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"and": schema.MapAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"not": schema.MapAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"or": schema.MapAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"regexp": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"flags": schema.ListAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"pattern": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"template": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"type": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"value": schema.StringAttribute{
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"field": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
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

					"global_output_refs": schema.ListAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"local_output_refs": schema.ListAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"logging_ref": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"match": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"and": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"not": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"or": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"regexp": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"flags": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"pattern": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"template": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"type": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"value": schema.StringAttribute{
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

func (r *LoggingBanzaicloudIoSyslogNgflowV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_logging_banzaicloud_io_syslog_ng_flow_v1beta1_manifest")

	var model LoggingBanzaicloudIoSyslogNgflowV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("logging.banzaicloud.io/v1beta1")
	model.Kind = pointer.String("SyslogNGFlow")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
