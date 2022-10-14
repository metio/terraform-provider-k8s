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

type ChartsFlagsmithComFlagsmithV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*ChartsFlagsmithComFlagsmithV1Alpha1Resource)(nil)
)

type ChartsFlagsmithComFlagsmithV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type ChartsFlagsmithComFlagsmithV1Alpha1GoModel struct {
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
		Api utilities.Dynamic `tfsdk:"api" yaml:"api,omitempty"`

		Frontend utilities.Dynamic `tfsdk:"frontend" yaml:"frontend,omitempty"`

		Hooks utilities.Dynamic `tfsdk:"hooks" yaml:"hooks,omitempty"`

		Influxdb *struct {
			Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`
		} `tfsdk:"influxdb" yaml:"influxdb,omitempty"`

		Ingress utilities.Dynamic `tfsdk:"ingress" yaml:"ingress,omitempty"`

		Metrics utilities.Dynamic `tfsdk:"metrics" yaml:"metrics,omitempty"`

		Openshift *bool `tfsdk:"openshift" yaml:"openshift,omitempty"`

		Postgresql *struct {
			Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`
		} `tfsdk:"postgresql" yaml:"postgresql,omitempty"`

		Service utilities.Dynamic `tfsdk:"service" yaml:"service,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewChartsFlagsmithComFlagsmithV1Alpha1Resource() resource.Resource {
	return &ChartsFlagsmithComFlagsmithV1Alpha1Resource{}
}

func (r *ChartsFlagsmithComFlagsmithV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_charts_flagsmith_com_flagsmith_v1alpha1"
}

func (r *ChartsFlagsmithComFlagsmithV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "Flagsmith is the Schema for the flagsmiths API",
		MarkdownDescription: "Flagsmith is the Schema for the flagsmiths API",
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
				Description:         "Spec defines the desired state of Flagsmith",
				MarkdownDescription: "Spec defines the desired state of Flagsmith",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"api": {
						Description:         "Configuration how to setup the flagsmith api service.",
						MarkdownDescription: "Configuration how to setup the flagsmith api service.",

						Type: utilities.DynamicType{},

						Required: true,
						Optional: false,
						Computed: false,
					},

					"frontend": {
						Description:         "Configuration how to setup the flagsmith frontend service.",
						MarkdownDescription: "Configuration how to setup the flagsmith frontend service.",

						Type: utilities.DynamicType{},

						Required: true,
						Optional: false,
						Computed: false,
					},

					"hooks": {
						Description:         "Configuration how to setup the flagsmith hooks.",
						MarkdownDescription: "Configuration how to setup the flagsmith hooks.",

						Type: utilities.DynamicType{},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"influxdb": {
						Description:         "Configuration how to setup the flagsmith InfluxDB service.",
						MarkdownDescription: "Configuration how to setup the flagsmith InfluxDB service.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"enabled": {
								Description:         "Set to true if InfluxDB will be installed. If the value is false InfluxDB will not be installed.",
								MarkdownDescription: "Set to true if InfluxDB will be installed. If the value is false InfluxDB will not be installed.",

								Type: types.BoolType,

								Required: true,
								Optional: false,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"ingress": {
						Description:         "Configuration how to setup ingress to the flagsmith if flagsmith is using Kubernetes and not OpenShift.",
						MarkdownDescription: "Configuration how to setup ingress to the flagsmith if flagsmith is using Kubernetes and not OpenShift.",

						Type: utilities.DynamicType{},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"metrics": {
						Description:         "Configuration how to setup the flagsmith metrics.",
						MarkdownDescription: "Configuration how to setup the flagsmith metrics.",

						Type: utilities.DynamicType{},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"openshift": {
						Description:         "If flagsmith install on OpenShift set value to true otherwise false.",
						MarkdownDescription: "If flagsmith install on OpenShift set value to true otherwise false.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"postgresql": {
						Description:         "Configuration how to setup the flagsmith postgresql service.",
						MarkdownDescription: "Configuration how to setup the flagsmith postgresql service.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"enabled": {
								Description:         "Set to true if PostgreSQL will be installed. If the value is false PostgreSQL will not be installed.",
								MarkdownDescription: "Set to true if PostgreSQL will be installed. If the value is false PostgreSQL will not be installed.",

								Type: types.BoolType,

								Required: true,
								Optional: false,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"service": {
						Description:         "Configuration how to setup the flagsmith kubernetes service.",
						MarkdownDescription: "Configuration how to setup the flagsmith kubernetes service.",

						Type: utilities.DynamicType{},

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

func (r *ChartsFlagsmithComFlagsmithV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_charts_flagsmith_com_flagsmith_v1alpha1")

	var state ChartsFlagsmithComFlagsmithV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel ChartsFlagsmithComFlagsmithV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("charts.flagsmith.com/v1alpha1")
	goModel.Kind = utilities.Ptr("Flagsmith")

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

func (r *ChartsFlagsmithComFlagsmithV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_charts_flagsmith_com_flagsmith_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *ChartsFlagsmithComFlagsmithV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_charts_flagsmith_com_flagsmith_v1alpha1")

	var state ChartsFlagsmithComFlagsmithV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel ChartsFlagsmithComFlagsmithV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("charts.flagsmith.com/v1alpha1")
	goModel.Kind = utilities.Ptr("Flagsmith")

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

func (r *ChartsFlagsmithComFlagsmithV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_charts_flagsmith_com_flagsmith_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
