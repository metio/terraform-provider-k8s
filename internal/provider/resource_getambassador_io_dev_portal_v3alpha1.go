/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"

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

type GetambassadorIoDevPortalV3Alpha1Resource struct{}

var (
	_ resource.Resource = (*GetambassadorIoDevPortalV3Alpha1Resource)(nil)
)

type GetambassadorIoDevPortalV3Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type GetambassadorIoDevPortalV3Alpha1GoModel struct {
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
		Ambassador_id *[]string `tfsdk:"ambassador_id" yaml:"ambassador_id,omitempty"`

		Content *struct {
			Branch *string `tfsdk:"branch" yaml:"branch,omitempty"`

			Dir *string `tfsdk:"dir" yaml:"dir,omitempty"`

			Url *string `tfsdk:"url" yaml:"url,omitempty"`
		} `tfsdk:"content" yaml:"content,omitempty"`

		Default *bool `tfsdk:"default" yaml:"default,omitempty"`

		Docs *[]struct {
			Service *string `tfsdk:"service" yaml:"service,omitempty"`

			Timeout_ms *int64 `tfsdk:"timeout_ms" yaml:"timeout_ms,omitempty"`

			Url *string `tfsdk:"url" yaml:"url,omitempty"`
		} `tfsdk:"docs" yaml:"docs,omitempty"`

		Naming_scheme *string `tfsdk:"naming_scheme" yaml:"naming_scheme,omitempty"`

		Preserve_servers *bool `tfsdk:"preserve_servers" yaml:"preserve_servers,omitempty"`

		Search *struct {
			Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

			Type *string `tfsdk:"type" yaml:"type,omitempty"`
		} `tfsdk:"search" yaml:"search,omitempty"`

		Selector *struct {
			MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`

			MatchNamespaces *[]string `tfsdk:"match_namespaces" yaml:"matchNamespaces,omitempty"`
		} `tfsdk:"selector" yaml:"selector,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewGetambassadorIoDevPortalV3Alpha1Resource() resource.Resource {
	return &GetambassadorIoDevPortalV3Alpha1Resource{}
}

func (r *GetambassadorIoDevPortalV3Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_getambassador_io_dev_portal_v3alpha1"
}

func (r *GetambassadorIoDevPortalV3Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "DevPortal is the Schema for the DevPortals API  DevPortal resources specify the 'what' and 'how' is shown in a DevPortal:   1. 'what' is in a DevPortal can be controlled with      - a 'selector', that can be used for filtering 'Mappings'.      - a 'docs' listing of (services, url)   2. 'how' is a pointer to some 'contents' (a checkout of a Git repository     with go-templates/markdown/css).  Multiple 'DevPortal's can exist in the cluster, and the Dev Portal server will show them at different endpoints. A 'DevPortal' resource with a special name, 'ambassador', will be used for configuring the default Dev Portal (served at '/docs/' by default).",
		MarkdownDescription: "DevPortal is the Schema for the DevPortals API  DevPortal resources specify the 'what' and 'how' is shown in a DevPortal:   1. 'what' is in a DevPortal can be controlled with      - a 'selector', that can be used for filtering 'Mappings'.      - a 'docs' listing of (services, url)   2. 'how' is a pointer to some 'contents' (a checkout of a Git repository     with go-templates/markdown/css).  Multiple 'DevPortal's can exist in the cluster, and the Dev Portal server will show them at different endpoints. A 'DevPortal' resource with a special name, 'ambassador', will be used for configuring the default Dev Portal (served at '/docs/' by default).",
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
				Description:         "DevPortalSpec defines the desired state of DevPortal",
				MarkdownDescription: "DevPortalSpec defines the desired state of DevPortal",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"ambassador_id": {
						Description:         "AmbassadorID declares which Ambassador instances should pay attention to this resource. If no value is provided, the default is:  	ambassador_id: 	- 'default'  TODO(lukeshu): In v3alpha2, consider renaming all of the 'ambassador_id' (singular) fields to 'ambassador_ids' (plural).",
						MarkdownDescription: "AmbassadorID declares which Ambassador instances should pay attention to this resource. If no value is provided, the default is:  	ambassador_id: 	- 'default'  TODO(lukeshu): In v3alpha2, consider renaming all of the 'ambassador_id' (singular) fields to 'ambassador_ids' (plural).",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"content": {
						Description:         "Content specifies where the content shown in the DevPortal come from",
						MarkdownDescription: "Content specifies where the content shown in the DevPortal come from",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"branch": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"dir": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"url": {
								Description:         "",
								MarkdownDescription: "",

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

					"default": {
						Description:         "Default must be true when this is the default DevPortal",
						MarkdownDescription: "Default must be true when this is the default DevPortal",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"docs": {
						Description:         "Docs is a static docs definition",
						MarkdownDescription: "Docs is a static docs definition",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"service": {
								Description:         "Service is the service being documented",
								MarkdownDescription: "Service is the service being documented",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"timeout_ms": {
								Description:         "Timeout specifies the amount of time devportal will wait for the downstream service to report an openapi spec back",
								MarkdownDescription: "Timeout specifies the amount of time devportal will wait for the downstream service to report an openapi spec back",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"url": {
								Description:         "URL is the URL used for obtaining docs",
								MarkdownDescription: "URL is the URL used for obtaining docs",

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

					"naming_scheme": {
						Description:         "Describes how to display 'services' in the DevPortal. Default namespace.name",
						MarkdownDescription: "Describes how to display 'services' in the DevPortal. Default namespace.name",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.OneOf("namespace.name", "name.prefix"),
						},
					},

					"preserve_servers": {
						Description:         "Configures this DevPortal to use server definitions from the openAPI doc instead of rewriting them based on the url used for the connection.",
						MarkdownDescription: "Configures this DevPortal to use server definitions from the openAPI doc instead of rewriting them based on the url used for the connection.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"search": {
						Description:         "DevPortalSearchSpec allows configuration over search functionality for the DevPortal",
						MarkdownDescription: "DevPortalSearchSpec allows configuration over search functionality for the DevPortal",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"enabled": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"type": {
								Description:         "Type of search. 'title-only' does a fuzzy search over openapi and page titles 'all-content' will fuzzy search over all openapi and page content. 'title-only' is the default. warning:  using all-content may incur a larger memory footprint",
								MarkdownDescription: "Type of search. 'title-only' does a fuzzy search over openapi and page titles 'all-content' will fuzzy search over all openapi and page content. 'title-only' is the default. warning:  using all-content may incur a larger memory footprint",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.OneOf("title-only", "all-content"),
								},
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"selector": {
						Description:         "Selector is used for choosing what is shown in the DevPortal",
						MarkdownDescription: "Selector is used for choosing what is shown in the DevPortal",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"match_labels": {
								Description:         "MatchLabels specifies the list of labels that must be present in Mappings for being present in this DevPortal.",
								MarkdownDescription: "MatchLabels specifies the list of labels that must be present in Mappings for being present in this DevPortal.",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"match_namespaces": {
								Description:         "MatchNamespaces is a list of namespaces that will be included in this DevPortal.",
								MarkdownDescription: "MatchNamespaces is a list of namespaces that will be included in this DevPortal.",

								Type: types.ListType{ElemType: types.StringType},

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
		},
	}, nil
}

func (r *GetambassadorIoDevPortalV3Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_getambassador_io_dev_portal_v3alpha1")

	var state GetambassadorIoDevPortalV3Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel GetambassadorIoDevPortalV3Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("getambassador.io/v3alpha1")
	goModel.Kind = utilities.Ptr("DevPortal")

	state.Id = types.Int64{Value: time.Now().UnixNano()}
	state.ApiVersion = types.String{Value: *goModel.ApiVersion}
	state.Kind = types.String{Value: *goModel.Kind}

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.String{Value: string(marshal)}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *GetambassadorIoDevPortalV3Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_getambassador_io_dev_portal_v3alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *GetambassadorIoDevPortalV3Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_getambassador_io_dev_portal_v3alpha1")

	var state GetambassadorIoDevPortalV3Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel GetambassadorIoDevPortalV3Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("getambassador.io/v3alpha1")
	goModel.Kind = utilities.Ptr("DevPortal")

	state.Id = types.Int64{Value: time.Now().UnixNano()}
	state.ApiVersion = types.String{Value: *goModel.ApiVersion}
	state.Kind = types.String{Value: *goModel.Kind}

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.String{Value: string(marshal)}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *GetambassadorIoDevPortalV3Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_getambassador_io_dev_portal_v3alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
