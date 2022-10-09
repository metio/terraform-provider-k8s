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

type IntegreatlyOrgGrafanaDashboardV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*IntegreatlyOrgGrafanaDashboardV1Alpha1Resource)(nil)
)

type IntegreatlyOrgGrafanaDashboardV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type IntegreatlyOrgGrafanaDashboardV1Alpha1GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Namespace *string `tfsdk:"namespace" yaml:"namespace"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		ConfigMapRef *struct {
			Key *string `tfsdk:"key" yaml:"key,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
		} `tfsdk:"config_map_ref" yaml:"configMapRef,omitempty"`

		ContentCacheDuration *string `tfsdk:"content_cache_duration" yaml:"contentCacheDuration,omitempty"`

		CustomFolderName *string `tfsdk:"custom_folder_name" yaml:"customFolderName,omitempty"`

		Datasources *[]struct {
			DatasourceName *string `tfsdk:"datasource_name" yaml:"datasourceName,omitempty"`

			InputName *string `tfsdk:"input_name" yaml:"inputName,omitempty"`
		} `tfsdk:"datasources" yaml:"datasources,omitempty"`

		GrafanaCom *struct {
			Id *int64 `tfsdk:"id" yaml:"id,omitempty"`

			Revision *int64 `tfsdk:"revision" yaml:"revision,omitempty"`
		} `tfsdk:"grafana_com" yaml:"grafanaCom,omitempty"`

		GzipConfigMapRef *struct {
			Key *string `tfsdk:"key" yaml:"key,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
		} `tfsdk:"gzip_config_map_ref" yaml:"gzipConfigMapRef,omitempty"`

		GzipJson *string `tfsdk:"gzip_json" yaml:"gzipJson,omitempty"`

		Json *string `tfsdk:"json" yaml:"json,omitempty"`

		Jsonnet *string `tfsdk:"jsonnet" yaml:"jsonnet,omitempty"`

		Plugins *[]struct {
			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Version *string `tfsdk:"version" yaml:"version,omitempty"`
		} `tfsdk:"plugins" yaml:"plugins,omitempty"`

		Url *string `tfsdk:"url" yaml:"url,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewIntegreatlyOrgGrafanaDashboardV1Alpha1Resource() resource.Resource {
	return &IntegreatlyOrgGrafanaDashboardV1Alpha1Resource{}
}

func (r *IntegreatlyOrgGrafanaDashboardV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_integreatly_org_grafana_dashboard_v1alpha1"
}

func (r *IntegreatlyOrgGrafanaDashboardV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "GrafanaDashboard is the Schema for the grafanadashboards API",
		MarkdownDescription: "GrafanaDashboard is the Schema for the grafanadashboards API",
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
				Description:         "GrafanaDashboardSpec defines the desired state of GrafanaDashboard",
				MarkdownDescription: "GrafanaDashboardSpec defines the desired state of GrafanaDashboard",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"config_map_ref": {
						Description:         "ConfigMapRef is a reference to a ConfigMap data field containing the dashboard's JSON",
						MarkdownDescription: "ConfigMapRef is a reference to a ConfigMap data field containing the dashboard's JSON",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"key": {
								Description:         "The key to select.",
								MarkdownDescription: "The key to select.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"name": {
								Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
								MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"optional": {
								Description:         "Specify whether the ConfigMap or its key must be defined",
								MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"content_cache_duration": {
						Description:         "ContentCacheDuration sets how often the operator should resync with the external source when using the 'grafanaCom.id' or 'url' field to specify the source of the dashboard. The default value is decided by the 'dashboardContentCacheDuration' field in the 'Grafana' resource. The default is 0 which is interpreted as never refetching.",
						MarkdownDescription: "ContentCacheDuration sets how often the operator should resync with the external source when using the 'grafanaCom.id' or 'url' field to specify the source of the dashboard. The default value is decided by the 'dashboardContentCacheDuration' field in the 'Grafana' resource. The default is 0 which is interpreted as never refetching.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"custom_folder_name": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"datasources": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"datasource_name": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"input_name": {
								Description:         "",
								MarkdownDescription: "",

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

					"grafana_com": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"id": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"revision": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"gzip_config_map_ref": {
						Description:         "GzipConfigMapRef is a reference to a ConfigMap binaryData field containing the dashboard's JSON, compressed with Gzip.",
						MarkdownDescription: "GzipConfigMapRef is a reference to a ConfigMap binaryData field containing the dashboard's JSON, compressed with Gzip.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"key": {
								Description:         "The key to select.",
								MarkdownDescription: "The key to select.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"name": {
								Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
								MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"optional": {
								Description:         "Specify whether the ConfigMap or its key must be defined",
								MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"gzip_json": {
						Description:         "GzipJson the dashboard's JSON compressed with Gzip. Base64-encoded when in YAML.",
						MarkdownDescription: "GzipJson the dashboard's JSON compressed with Gzip. Base64-encoded when in YAML.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"json": {
						Description:         "Json is the dashboard's JSON",
						MarkdownDescription: "Json is the dashboard's JSON",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"jsonnet": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"plugins": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"name": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"version": {
								Description:         "",
								MarkdownDescription: "",

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
		},
	}, nil
}

func (r *IntegreatlyOrgGrafanaDashboardV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_integreatly_org_grafana_dashboard_v1alpha1")

	var state IntegreatlyOrgGrafanaDashboardV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel IntegreatlyOrgGrafanaDashboardV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("integreatly.org/v1alpha1")
	goModel.Kind = utilities.Ptr("GrafanaDashboard")

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

func (r *IntegreatlyOrgGrafanaDashboardV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_integreatly_org_grafana_dashboard_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *IntegreatlyOrgGrafanaDashboardV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_integreatly_org_grafana_dashboard_v1alpha1")

	var state IntegreatlyOrgGrafanaDashboardV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel IntegreatlyOrgGrafanaDashboardV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("integreatly.org/v1alpha1")
	goModel.Kind = utilities.Ptr("GrafanaDashboard")

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

func (r *IntegreatlyOrgGrafanaDashboardV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_integreatly_org_grafana_dashboard_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
