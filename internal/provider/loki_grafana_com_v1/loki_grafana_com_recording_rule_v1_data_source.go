/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package loki_grafana_com_v1

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
	_ datasource.DataSource              = &LokiGrafanaComRecordingRuleV1DataSource{}
	_ datasource.DataSourceWithConfigure = &LokiGrafanaComRecordingRuleV1DataSource{}
)

func NewLokiGrafanaComRecordingRuleV1DataSource() datasource.DataSource {
	return &LokiGrafanaComRecordingRuleV1DataSource{}
}

type LokiGrafanaComRecordingRuleV1DataSource struct {
	kubernetesClient dynamic.Interface
}

type LokiGrafanaComRecordingRuleV1DataSourceData struct {
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
		Groups *[]struct {
			Interval *string `tfsdk:"interval" json:"interval,omitempty"`
			Limit    *int64  `tfsdk:"limit" json:"limit,omitempty"`
			Name     *string `tfsdk:"name" json:"name,omitempty"`
			Rules    *[]struct {
				Expr   *string            `tfsdk:"expr" json:"expr,omitempty"`
				Labels *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
				Record *string            `tfsdk:"record" json:"record,omitempty"`
			} `tfsdk:"rules" json:"rules,omitempty"`
		} `tfsdk:"groups" json:"groups,omitempty"`
		TenantID *string `tfsdk:"tenant_id" json:"tenantID,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *LokiGrafanaComRecordingRuleV1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_loki_grafana_com_recording_rule_v1"
}

func (r *LokiGrafanaComRecordingRuleV1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "RecordingRule is the Schema for the recordingrules API",
		MarkdownDescription: "RecordingRule is the Schema for the recordingrules API",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"api_version": schema.StringAttribute{
				Description:         "The API group of the requested resource.",
				MarkdownDescription: "The API group of the requested resource.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"kind": schema.StringAttribute{
				Description:         "The type of the requested resource.",
				MarkdownDescription: "The type of the requested resource.",
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
				Description:         "RecordingRuleSpec defines the desired state of RecordingRule",
				MarkdownDescription: "RecordingRuleSpec defines the desired state of RecordingRule",
				Attributes: map[string]schema.Attribute{
					"groups": schema.ListNestedAttribute{
						Description:         "List of groups for recording rules.",
						MarkdownDescription: "List of groups for recording rules.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"interval": schema.StringAttribute{
									Description:         "Interval defines the time interval between evaluation of the given recoding rule.",
									MarkdownDescription: "Interval defines the time interval between evaluation of the given recoding rule.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"limit": schema.Int64Attribute{
									Description:         "Limit defines the number of series a recording rule can produce. 0 is no limit.",
									MarkdownDescription: "Limit defines the number of series a recording rule can produce. 0 is no limit.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"name": schema.StringAttribute{
									Description:         "Name of the recording rule group. Must be unique within all recording rules.",
									MarkdownDescription: "Name of the recording rule group. Must be unique within all recording rules.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"rules": schema.ListNestedAttribute{
									Description:         "Rules defines a list of recording rules",
									MarkdownDescription: "Rules defines a list of recording rules",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"expr": schema.StringAttribute{
												Description:         "The LogQL expression to evaluate. Every evaluation cycle this is evaluated at the current time, and all resultant time series become pending/firing alerts.",
												MarkdownDescription: "The LogQL expression to evaluate. Every evaluation cycle this is evaluated at the current time, and all resultant time series become pending/firing alerts.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"labels": schema.MapAttribute{
												Description:         "Labels to add to each recording rule.",
												MarkdownDescription: "Labels to add to each recording rule.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"record": schema.StringAttribute{
												Description:         "The name of the time series to output to. Must be a valid metric name.",
												MarkdownDescription: "The name of the time series to output to. Must be a valid metric name.",
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
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"tenant_id": schema.StringAttribute{
						Description:         "TenantID of tenant where the recording rules are evaluated in.",
						MarkdownDescription: "TenantID of tenant where the recording rules are evaluated in.",
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

func (r *LokiGrafanaComRecordingRuleV1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.Append(utilities.OfflineProviderError())
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.Append(utilities.UnexpectedDataSourceDataError(request.ProviderData))
	}
}

func (r *LokiGrafanaComRecordingRuleV1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_loki_grafana_com_recording_rule_v1")

	var data LokiGrafanaComRecordingRuleV1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "loki.grafana.com", Version: "v1", Resource: "recordingrules"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.Append(utilities.GetNamespacedResourceError(err, data.Metadata.Name, data.Metadata.Namespace))
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse LokiGrafanaComRecordingRuleV1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Namespace, data.Metadata.Name))
	data.ApiVersion = pointer.String("loki.grafana.com/v1")
	data.Kind = pointer.String("RecordingRule")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
