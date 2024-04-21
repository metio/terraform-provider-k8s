/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package canaries_flanksource_com_v1

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
	_ datasource.DataSource = &CanariesFlanksourceComComponentV1Manifest{}
)

func NewCanariesFlanksourceComComponentV1Manifest() datasource.DataSource {
	return &CanariesFlanksourceComComponentV1Manifest{}
}

type CanariesFlanksourceComComponentV1Manifest struct{}

type CanariesFlanksourceComComponentV1ManifestData struct {
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
		Checks *[]struct {
		} `tfsdk:"checks" json:"checks,omitempty"`
		Components *map[string]string `tfsdk:"components" json:"components,omitempty"`
		Configs    *[]struct {
			Agent         *string            `tfsdk:"agent" json:"agent,omitempty"`
			Cache         *string            `tfsdk:"cache" json:"cache,omitempty"`
			Class         *string            `tfsdk:"class" json:"class,omitempty"`
			External_id   *string            `tfsdk:"external_id" json:"external_id,omitempty"`
			FieldSelector *string            `tfsdk:"field_selector" json:"fieldSelector,omitempty"`
			Id            *string            `tfsdk:"id" json:"id,omitempty"`
			LabelSelector *string            `tfsdk:"label_selector" json:"labelSelector,omitempty"`
			Name          *string            `tfsdk:"name" json:"name,omitempty"`
			Namespace     *string            `tfsdk:"namespace" json:"namespace,omitempty"`
			Statuses      *[]string          `tfsdk:"statuses" json:"statuses,omitempty"`
			TagSelector   *string            `tfsdk:"tag_selector" json:"tagSelector,omitempty"`
			Tags          *map[string]string `tfsdk:"tags" json:"tags,omitempty"`
			Type          *string            `tfsdk:"type" json:"type,omitempty"`
			Types         *[]string          `tfsdk:"types" json:"types,omitempty"`
		} `tfsdk:"configs" json:"configs,omitempty"`
		ForEach *map[string]string `tfsdk:"for_each" json:"forEach,omitempty"`
		Hidden  *bool              `tfsdk:"hidden" json:"hidden,omitempty"`
		Icon    *string            `tfsdk:"icon" json:"icon,omitempty"`
		Id      *struct {
			Expr       *string `tfsdk:"expr" json:"expr,omitempty"`
			Javascript *string `tfsdk:"javascript" json:"javascript,omitempty"`
			JsonPath   *string `tfsdk:"json_path" json:"jsonPath,omitempty"`
			Template   *string `tfsdk:"template" json:"template,omitempty"`
		} `tfsdk:"id" json:"id,omitempty"`
		Labels    *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Lifecycle *string            `tfsdk:"lifecycle" json:"lifecycle,omitempty"`
		Logs      *[]struct {
			Labels *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			Name   *string            `tfsdk:"name" json:"name,omitempty"`
			Type   *string            `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"logs" json:"logs,omitempty"`
		Lookup       *map[string]string `tfsdk:"lookup" json:"lookup,omitempty"`
		Name         *string            `tfsdk:"name" json:"name,omitempty"`
		Namespace    *string            `tfsdk:"namespace" json:"namespace,omitempty"`
		Order        *int64             `tfsdk:"order" json:"order,omitempty"`
		Owner        *string            `tfsdk:"owner" json:"owner,omitempty"`
		ParentLookup *struct {
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			Type      *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"parent_lookup" json:"parentLookup,omitempty"`
		Properties    *map[string]string `tfsdk:"properties" json:"properties,omitempty"`
		Relationships *[]struct {
			Ref  *string `tfsdk:"ref" json:"ref,omitempty"`
			Type *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"relationships" json:"relationships,omitempty"`
		Selectors *[]struct {
			Agent         *string   `tfsdk:"agent" json:"agent,omitempty"`
			Cache         *string   `tfsdk:"cache" json:"cache,omitempty"`
			FieldSelector *string   `tfsdk:"field_selector" json:"fieldSelector,omitempty"`
			Id            *string   `tfsdk:"id" json:"id,omitempty"`
			LabelSelector *string   `tfsdk:"label_selector" json:"labelSelector,omitempty"`
			Name          *string   `tfsdk:"name" json:"name,omitempty"`
			Namespace     *string   `tfsdk:"namespace" json:"namespace,omitempty"`
			Statuses      *[]string `tfsdk:"statuses" json:"statuses,omitempty"`
			TagSelector   *string   `tfsdk:"tag_selector" json:"tagSelector,omitempty"`
			Types         *[]string `tfsdk:"types" json:"types,omitempty"`
		} `tfsdk:"selectors" json:"selectors,omitempty"`
		Summary *struct {
			Checks    *map[string]string `tfsdk:"checks" json:"checks,omitempty"`
			Healthy   *int64             `tfsdk:"healthy" json:"healthy,omitempty"`
			Incidents *struct {
			} `tfsdk:"incidents" json:"incidents,omitempty"`
			Info     *int64 `tfsdk:"info" json:"info,omitempty"`
			Insights *struct {
			} `tfsdk:"insights" json:"insights,omitempty"`
			Unhealthy *int64 `tfsdk:"unhealthy" json:"unhealthy,omitempty"`
			Warning   *int64 `tfsdk:"warning" json:"warning,omitempty"`
		} `tfsdk:"summary" json:"summary,omitempty"`
		Tooltip *string `tfsdk:"tooltip" json:"tooltip,omitempty"`
		Type    *string `tfsdk:"type" json:"type,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CanariesFlanksourceComComponentV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_canaries_flanksource_com_component_v1_manifest"
}

func (r *CanariesFlanksourceComComponentV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "",
		MarkdownDescription: "",
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
					"checks": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"components": schema.MapAttribute{
						Description:         "Create new child components",
						MarkdownDescription: "Create new child components",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"configs": schema.ListNestedAttribute{
						Description:         "Lookup and associate config items with this component",
						MarkdownDescription: "Lookup and associate config items with this component",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"agent": schema.StringAttribute{
									Description:         "Agent can be the agent id or the name of the agent. Additionally, the special 'self' value can be used to select resources without an agent.",
									MarkdownDescription: "Agent can be the agent id or the name of the agent. Additionally, the special 'self' value can be used to select resources without an agent.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"cache": schema.StringAttribute{
									Description:         "Cache directives 'no-cache' (should not fetch from cache but can be cached) 'no-store' (should not cache) 'max-age=X' (cache for X duration)",
									MarkdownDescription: "Cache directives 'no-cache' (should not fetch from cache but can be cached) 'no-store' (should not cache) 'max-age=X' (cache for X duration)",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"class": schema.StringAttribute{
									Description:         "Deprecated. Use 'fieldSelector (config_class=)'",
									MarkdownDescription: "Deprecated. Use 'fieldSelector (config_class=)'",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"external_id": schema.StringAttribute{
									Description:         "Deprecated. Use 'fieldSelector (external_id=)'",
									MarkdownDescription: "Deprecated. Use 'fieldSelector (external_id=)'",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"field_selector": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"id": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"label_selector": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"namespace": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"statuses": schema.ListAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"tag_selector": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"tags": schema.MapAttribute{
									Description:         "Deprecated. Use 'labelSelector'",
									MarkdownDescription: "Deprecated. Use 'labelSelector'",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"type": schema.StringAttribute{
									Description:         "Deprecated. Use 'types'",
									MarkdownDescription: "Deprecated. Use 'types'",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"types": schema.ListAttribute{
									Description:         "",
									MarkdownDescription: "",
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

					"for_each": schema.MapAttribute{
						Description:         "Only applies when using lookup, when specified the components and propertiesspecified under ForEach will be templated using the components returned by the lookup${.properties} can be used to reference the properties of the component${.component} can be used to reference the component itself",
						MarkdownDescription: "Only applies when using lookup, when specified the components and propertiesspecified under ForEach will be templated using the components returned by the lookup${.properties} can be used to reference the properties of the component${.component} can be used to reference the component itself",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"hidden": schema.BoolAttribute{
						Description:         "If set to true, do not display in UI",
						MarkdownDescription: "If set to true, do not display in UI",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"icon": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"id": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"expr": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"javascript": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"json_path": schema.StringAttribute{
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

					"labels": schema.MapAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"lifecycle": schema.StringAttribute{
						Description:         "The lifecycle state of the component e.g. production, staging, dev, etc.",
						MarkdownDescription: "The lifecycle state of the component e.g. production, staging, dev, etc.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"logs": schema.ListNestedAttribute{
						Description:         "Logs is a list of logs selector for apm-hub.",
						MarkdownDescription: "Logs is a list of logs selector for apm-hub.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"labels": schema.MapAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"name": schema.StringAttribute{
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
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"lookup": schema.MapAttribute{
						Description:         "Lookup component definitions from an external source, use theforEach property to iterate over the results to further enrich each component.",
						MarkdownDescription: "Lookup component definitions from an external source, use theforEach property to iterate over the results to further enrich each component.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"name": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"namespace": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"order": schema.Int64Attribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"owner": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"parent_lookup": schema.SingleNestedAttribute{
						Description:         "Reference to populate parent_id",
						MarkdownDescription: "Reference to populate parent_id",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"namespace": schema.StringAttribute{
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"properties": schema.MapAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"relationships": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"ref": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"type": schema.StringAttribute{
									Description:         "The type of relationship, e.g. dependsOn, subcomponentOf, providesApis, consumesApis",
									MarkdownDescription: "The type of relationship, e.g. dependsOn, subcomponentOf, providesApis, consumesApis",
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

					"selectors": schema.ListNestedAttribute{
						Description:         "Lookup and associcate other components with this component",
						MarkdownDescription: "Lookup and associcate other components with this component",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"agent": schema.StringAttribute{
									Description:         "Agent can be the agent id or the name of the agent. Additionally, the special 'self' value can be used to select resources without an agent.",
									MarkdownDescription: "Agent can be the agent id or the name of the agent. Additionally, the special 'self' value can be used to select resources without an agent.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"cache": schema.StringAttribute{
									Description:         "Cache directives 'no-cache' (should not fetch from cache but can be cached) 'no-store' (should not cache) 'max-age=X' (cache for X duration)",
									MarkdownDescription: "Cache directives 'no-cache' (should not fetch from cache but can be cached) 'no-store' (should not cache) 'max-age=X' (cache for X duration)",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"field_selector": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"id": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"label_selector": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"namespace": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"statuses": schema.ListAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"tag_selector": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"types": schema.ListAttribute{
									Description:         "",
									MarkdownDescription: "",
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

					"summary": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"checks": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"healthy": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"incidents": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes:          map[string]schema.Attribute{},
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"info": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"insights": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes:          map[string]schema.Attribute{},
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"unhealthy": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"warning": schema.Int64Attribute{
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

					"tooltip": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"type": schema.StringAttribute{
						Description:         "The type of component, e.g. service, API, website, library, database, etc.",
						MarkdownDescription: "The type of component, e.g. service, API, website, library, database, etc.",
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

func (r *CanariesFlanksourceComComponentV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_canaries_flanksource_com_component_v1_manifest")

	var model CanariesFlanksourceComComponentV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("canaries.flanksource.com/v1")
	model.Kind = pointer.String("Component")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
