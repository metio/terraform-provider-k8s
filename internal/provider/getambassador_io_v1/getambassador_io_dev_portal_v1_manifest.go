/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package getambassador_io_v1

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
	_ datasource.DataSource = &GetambassadorIoDevPortalV1Manifest{}
)

func NewGetambassadorIoDevPortalV1Manifest() datasource.DataSource {
	return &GetambassadorIoDevPortalV1Manifest{}
}

type GetambassadorIoDevPortalV1Manifest struct{}

type GetambassadorIoDevPortalV1ManifestData struct {
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
		Ambassador_id *[]string `tfsdk:"ambassador_id" json:"ambassador_id,omitempty"`
		Content       *struct {
			Branch *string `tfsdk:"branch" json:"branch,omitempty"`
			Dir    *string `tfsdk:"dir" json:"dir,omitempty"`
			Url    *string `tfsdk:"url" json:"url,omitempty"`
		} `tfsdk:"content" json:"content,omitempty"`
		Default *bool `tfsdk:"default" json:"default,omitempty"`
		Docs    *[]struct {
			Service    *string `tfsdk:"service" json:"service,omitempty"`
			Timeout_ms *int64  `tfsdk:"timeout_ms" json:"timeout_ms,omitempty"`
			Url        *string `tfsdk:"url" json:"url,omitempty"`
		} `tfsdk:"docs" json:"docs,omitempty"`
		Naming_scheme    *string `tfsdk:"naming_scheme" json:"naming_scheme,omitempty"`
		Preserve_servers *bool   `tfsdk:"preserve_servers" json:"preserve_servers,omitempty"`
		Search           *struct {
			Enabled *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
			Type    *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"search" json:"search,omitempty"`
		Selector *struct {
			MatchLabels     *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			MatchNamespaces *[]string          `tfsdk:"match_namespaces" json:"matchNamespaces,omitempty"`
		} `tfsdk:"selector" json:"selector,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *GetambassadorIoDevPortalV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_getambassador_io_dev_portal_v1_manifest"
}

func (r *GetambassadorIoDevPortalV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "DevPortal is the Schema for the DevPortals API DevPortal resources specify the 'what' and 'how' is shown in a DevPortal: 1. 'what' is in a DevPortal can be controlled with - a 'selector', that can be used for filtering 'Mappings'. - a 'docs' listing of (services, url) 2. 'how' is a pointer to some 'contents' (a checkout of a Git repository with go-templates/markdown/css). Multiple 'DevPortal's can exist in the cluster, and the Dev Portal server will show them at different endpoints. A 'DevPortal' resource with a special name, 'ambassador', will be used for configuring the default Dev Portal (served at '/docs/' by default).",
		MarkdownDescription: "DevPortal is the Schema for the DevPortals API DevPortal resources specify the 'what' and 'how' is shown in a DevPortal: 1. 'what' is in a DevPortal can be controlled with - a 'selector', that can be used for filtering 'Mappings'. - a 'docs' listing of (services, url) 2. 'how' is a pointer to some 'contents' (a checkout of a Git repository with go-templates/markdown/css). Multiple 'DevPortal's can exist in the cluster, and the Dev Portal server will show them at different endpoints. A 'DevPortal' resource with a special name, 'ambassador', will be used for configuring the default Dev Portal (served at '/docs/' by default).",
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
				Description:         "DevPortalSpec defines the desired state of DevPortal",
				MarkdownDescription: "DevPortalSpec defines the desired state of DevPortal",
				Attributes: map[string]schema.Attribute{
					"ambassador_id": schema.ListAttribute{
						Description:         "AmbassadorID declares which Ambassador instances should pay attention to this resource. May either be a string or a list of strings. If no value is provided, the default is: ambassador_id: - 'default'",
						MarkdownDescription: "AmbassadorID declares which Ambassador instances should pay attention to this resource. May either be a string or a list of strings. If no value is provided, the default is: ambassador_id: - 'default'",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"content": schema.SingleNestedAttribute{
						Description:         "Content specifies where the content shown in the DevPortal come from",
						MarkdownDescription: "Content specifies where the content shown in the DevPortal come from",
						Attributes: map[string]schema.Attribute{
							"branch": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"dir": schema.StringAttribute{
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

					"default": schema.BoolAttribute{
						Description:         "Default must be true when this is the default DevPortal",
						MarkdownDescription: "Default must be true when this is the default DevPortal",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"docs": schema.ListNestedAttribute{
						Description:         "Docs is a static docs definition",
						MarkdownDescription: "Docs is a static docs definition",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"service": schema.StringAttribute{
									Description:         "Service is the service being documented",
									MarkdownDescription: "Service is the service being documented",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"timeout_ms": schema.Int64Attribute{
									Description:         "Timeout specifies the amount of time devportal will wait for the downstream service to report an openapi spec back",
									MarkdownDescription: "Timeout specifies the amount of time devportal will wait for the downstream service to report an openapi spec back",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"url": schema.StringAttribute{
									Description:         "URL is the URL used for obtaining docs",
									MarkdownDescription: "URL is the URL used for obtaining docs",
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

					"naming_scheme": schema.StringAttribute{
						Description:         "Describes how to display 'services' in the DevPortal. Default namespace.name",
						MarkdownDescription: "Describes how to display 'services' in the DevPortal. Default namespace.name",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("namespace.name", "name.prefix"),
						},
					},

					"preserve_servers": schema.BoolAttribute{
						Description:         "Configures this DevPortal to use server definitions from the openAPI doc instead of rewriting them based on the url used for the connection.",
						MarkdownDescription: "Configures this DevPortal to use server definitions from the openAPI doc instead of rewriting them based on the url used for the connection.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"search": schema.SingleNestedAttribute{
						Description:         "DevPortalSearchSpec allows configuration over search functionality for the DevPortal",
						MarkdownDescription: "DevPortalSearchSpec allows configuration over search functionality for the DevPortal",
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"type": schema.StringAttribute{
								Description:         "Type of search. 'title-only' does a fuzzy search over openapi and page titles 'all-content' will fuzzy search over all openapi and page content. 'title-only' is the default. warning: using all-content may incur a larger memory footprint",
								MarkdownDescription: "Type of search. 'title-only' does a fuzzy search over openapi and page titles 'all-content' will fuzzy search over all openapi and page content. 'title-only' is the default. warning: using all-content may incur a larger memory footprint",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("title-only", "all-content"),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"selector": schema.SingleNestedAttribute{
						Description:         "Selector is used for choosing what is shown in the DevPortal",
						MarkdownDescription: "Selector is used for choosing what is shown in the DevPortal",
						Attributes: map[string]schema.Attribute{
							"match_labels": schema.MapAttribute{
								Description:         "MatchLabels specifies the list of labels that must be present in Mappings for being present in this DevPortal.",
								MarkdownDescription: "MatchLabels specifies the list of labels that must be present in Mappings for being present in this DevPortal.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"match_namespaces": schema.ListAttribute{
								Description:         "MatchNamespaces is a list of namespaces that will be included in this DevPortal.",
								MarkdownDescription: "MatchNamespaces is a list of namespaces that will be included in this DevPortal.",
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

func (r *GetambassadorIoDevPortalV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_getambassador_io_dev_portal_v1_manifest")

	var model GetambassadorIoDevPortalV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("getambassador.io/v1")
	model.Kind = pointer.String("DevPortal")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
