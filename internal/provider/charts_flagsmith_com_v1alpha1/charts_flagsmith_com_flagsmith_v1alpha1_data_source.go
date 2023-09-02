/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package charts_flagsmith_com_v1alpha1

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
	_ datasource.DataSource              = &ChartsFlagsmithComFlagsmithV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &ChartsFlagsmithComFlagsmithV1Alpha1DataSource{}
)

func NewChartsFlagsmithComFlagsmithV1Alpha1DataSource() datasource.DataSource {
	return &ChartsFlagsmithComFlagsmithV1Alpha1DataSource{}
}

type ChartsFlagsmithComFlagsmithV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type ChartsFlagsmithComFlagsmithV1Alpha1DataSourceData struct {
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

func (r *ChartsFlagsmithComFlagsmithV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_charts_flagsmith_com_flagsmith_v1alpha1"
}

func (r *ChartsFlagsmithComFlagsmithV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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
				Description:         "Spec defines the desired state of Flagsmith",
				MarkdownDescription: "Spec defines the desired state of Flagsmith",
				Attributes: map[string]schema.Attribute{
					"api": schema.MapAttribute{
						Description:         "Configuration how to setup the flagsmith api service.",
						MarkdownDescription: "Configuration how to setup the flagsmith api service.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"frontend": schema.MapAttribute{
						Description:         "Configuration how to setup the flagsmith frontend service.",
						MarkdownDescription: "Configuration how to setup the flagsmith frontend service.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"hooks": schema.MapAttribute{
						Description:         "Configuration how to setup the flagsmith hooks.",
						MarkdownDescription: "Configuration how to setup the flagsmith hooks.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"influxdb": schema.SingleNestedAttribute{
						Description:         "Configuration how to setup the flagsmith InfluxDB service.",
						MarkdownDescription: "Configuration how to setup the flagsmith InfluxDB service.",
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description:         "Set to true if InfluxDB will be installed. If the value is false InfluxDB will not be installed.",
								MarkdownDescription: "Set to true if InfluxDB will be installed. If the value is false InfluxDB will not be installed.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"ingress": schema.MapAttribute{
						Description:         "Configuration how to setup ingress to the flagsmith if flagsmith is using Kubernetes and not OpenShift.",
						MarkdownDescription: "Configuration how to setup ingress to the flagsmith if flagsmith is using Kubernetes and not OpenShift.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"metrics": schema.MapAttribute{
						Description:         "Configuration how to setup the flagsmith metrics.",
						MarkdownDescription: "Configuration how to setup the flagsmith metrics.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"openshift": schema.BoolAttribute{
						Description:         "If flagsmith install on OpenShift set value to true otherwise false.",
						MarkdownDescription: "If flagsmith install on OpenShift set value to true otherwise false.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"postgresql": schema.SingleNestedAttribute{
						Description:         "Configuration how to setup the flagsmith postgresql service.",
						MarkdownDescription: "Configuration how to setup the flagsmith postgresql service.",
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description:         "Set to true if PostgreSQL will be installed. If the value is false PostgreSQL will not be installed.",
								MarkdownDescription: "Set to true if PostgreSQL will be installed. If the value is false PostgreSQL will not be installed.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"service": schema.MapAttribute{
						Description:         "Configuration how to setup the flagsmith kubernetes service.",
						MarkdownDescription: "Configuration how to setup the flagsmith kubernetes service.",
						ElementType:         types.StringType,
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
	}
}

func (r *ChartsFlagsmithComFlagsmithV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *ChartsFlagsmithComFlagsmithV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_charts_flagsmith_com_flagsmith_v1alpha1")

	var data ChartsFlagsmithComFlagsmithV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "charts.flagsmith.com", Version: "v1alpha1", Resource: "Flagsmith"}).
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

	var readResponse ChartsFlagsmithComFlagsmithV1Alpha1DataSourceData
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
	data.ApiVersion = pointer.String("charts.flagsmith.com/v1alpha1")
	data.Kind = pointer.String("Flagsmith")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}