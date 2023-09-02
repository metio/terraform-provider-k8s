/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package charts_flagsmith_com_v1alpha1

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &ChartsFlagsmithComFlagsmithV1Alpha1Manifest{}
)

func NewChartsFlagsmithComFlagsmithV1Alpha1Manifest() datasource.DataSource {
	return &ChartsFlagsmithComFlagsmithV1Alpha1Manifest{}
}

type ChartsFlagsmithComFlagsmithV1Alpha1Manifest struct{}

type ChartsFlagsmithComFlagsmithV1Alpha1ManifestData struct {
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
		Api      *map[string]string `tfsdk:"api" json:"api,omitempty"`
		Frontend *map[string]string `tfsdk:"frontend" json:"frontend,omitempty"`
		Hooks    *map[string]string `tfsdk:"hooks" json:"hooks,omitempty"`
		Influxdb *struct {
			Enabled *bool `tfsdk:"enabled" json:"enabled,omitempty"`
		} `tfsdk:"influxdb" json:"influxdb,omitempty"`
		Ingress    *map[string]string `tfsdk:"ingress" json:"ingress,omitempty"`
		Metrics    *map[string]string `tfsdk:"metrics" json:"metrics,omitempty"`
		Openshift  *bool              `tfsdk:"openshift" json:"openshift,omitempty"`
		Postgresql *struct {
			Enabled *bool `tfsdk:"enabled" json:"enabled,omitempty"`
		} `tfsdk:"postgresql" json:"postgresql,omitempty"`
		Service *map[string]string `tfsdk:"service" json:"service,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ChartsFlagsmithComFlagsmithV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_charts_flagsmith_com_flagsmith_v1alpha1_manifest"
}

func (r *ChartsFlagsmithComFlagsmithV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Flagsmith is the Schema for the flagsmiths API",
		MarkdownDescription: "Flagsmith is the Schema for the flagsmiths API",
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
				Description:         "Spec defines the desired state of Flagsmith",
				MarkdownDescription: "Spec defines the desired state of Flagsmith",
				Attributes: map[string]schema.Attribute{
					"api": schema.MapAttribute{
						Description:         "Configuration how to setup the flagsmith api service.",
						MarkdownDescription: "Configuration how to setup the flagsmith api service.",
						ElementType:         types.StringType,
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"frontend": schema.MapAttribute{
						Description:         "Configuration how to setup the flagsmith frontend service.",
						MarkdownDescription: "Configuration how to setup the flagsmith frontend service.",
						ElementType:         types.StringType,
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"hooks": schema.MapAttribute{
						Description:         "Configuration how to setup the flagsmith hooks.",
						MarkdownDescription: "Configuration how to setup the flagsmith hooks.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"influxdb": schema.SingleNestedAttribute{
						Description:         "Configuration how to setup the flagsmith InfluxDB service.",
						MarkdownDescription: "Configuration how to setup the flagsmith InfluxDB service.",
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description:         "Set to true if InfluxDB will be installed. If the value is false InfluxDB will not be installed.",
								MarkdownDescription: "Set to true if InfluxDB will be installed. If the value is false InfluxDB will not be installed.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"ingress": schema.MapAttribute{
						Description:         "Configuration how to setup ingress to the flagsmith if flagsmith is using Kubernetes and not OpenShift.",
						MarkdownDescription: "Configuration how to setup ingress to the flagsmith if flagsmith is using Kubernetes and not OpenShift.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"metrics": schema.MapAttribute{
						Description:         "Configuration how to setup the flagsmith metrics.",
						MarkdownDescription: "Configuration how to setup the flagsmith metrics.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"openshift": schema.BoolAttribute{
						Description:         "If flagsmith install on OpenShift set value to true otherwise false.",
						MarkdownDescription: "If flagsmith install on OpenShift set value to true otherwise false.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"postgresql": schema.SingleNestedAttribute{
						Description:         "Configuration how to setup the flagsmith postgresql service.",
						MarkdownDescription: "Configuration how to setup the flagsmith postgresql service.",
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description:         "Set to true if PostgreSQL will be installed. If the value is false PostgreSQL will not be installed.",
								MarkdownDescription: "Set to true if PostgreSQL will be installed. If the value is false PostgreSQL will not be installed.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"service": schema.MapAttribute{
						Description:         "Configuration how to setup the flagsmith kubernetes service.",
						MarkdownDescription: "Configuration how to setup the flagsmith kubernetes service.",
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
	}
}

func (r *ChartsFlagsmithComFlagsmithV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_charts_flagsmith_com_flagsmith_v1alpha1_manifest")

	var model ChartsFlagsmithComFlagsmithV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Name, model.Metadata.Namespace))
	model.ApiVersion = pointer.String("charts.flagsmith.com/v1alpha1")
	model.Kind = pointer.String("Flagsmith")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"YAML Error: "+err.Error(),
		)
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
