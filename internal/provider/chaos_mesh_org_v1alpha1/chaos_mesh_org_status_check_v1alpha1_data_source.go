/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package chaos_mesh_org_v1alpha1

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
	"net/http"
)

var (
	_ datasource.DataSource              = &ChaosMeshOrgStatusCheckV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &ChaosMeshOrgStatusCheckV1Alpha1DataSource{}
)

func NewChaosMeshOrgStatusCheckV1Alpha1DataSource() datasource.DataSource {
	return &ChaosMeshOrgStatusCheckV1Alpha1DataSource{}
}

type ChaosMeshOrgStatusCheckV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type ChaosMeshOrgStatusCheckV1Alpha1DataSourceData struct {
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
		Duration         *string `tfsdk:"duration" json:"duration,omitempty"`
		FailureThreshold *int64  `tfsdk:"failure_threshold" json:"failureThreshold,omitempty"`
		Http             *struct {
			Body     *string `tfsdk:"body" json:"body,omitempty"`
			Criteria *struct {
				StatusCode *string `tfsdk:"status_code" json:"statusCode,omitempty"`
			} `tfsdk:"criteria" json:"criteria,omitempty"`
			Headers *map[string][]string `tfsdk:"headers" json:"headers,omitempty"`
			Method  *string              `tfsdk:"method" json:"method,omitempty"`
			Url     *string              `tfsdk:"url" json:"url,omitempty"`
		} `tfsdk:"http" json:"http,omitempty"`
		IntervalSeconds     *int64  `tfsdk:"interval_seconds" json:"intervalSeconds,omitempty"`
		Mode                *string `tfsdk:"mode" json:"mode,omitempty"`
		RecordsHistoryLimit *int64  `tfsdk:"records_history_limit" json:"recordsHistoryLimit,omitempty"`
		SuccessThreshold    *int64  `tfsdk:"success_threshold" json:"successThreshold,omitempty"`
		TimeoutSeconds      *int64  `tfsdk:"timeout_seconds" json:"timeoutSeconds,omitempty"`
		Type                *string `tfsdk:"type" json:"type,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ChaosMeshOrgStatusCheckV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_chaos_mesh_org_status_check_v1alpha1"
}

func (r *ChaosMeshOrgStatusCheckV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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
				Description:         "Spec defines the behavior of a status check",
				MarkdownDescription: "Spec defines the behavior of a status check",
				Attributes: map[string]schema.Attribute{
					"duration": schema.StringAttribute{
						Description:         "Duration defines the duration of the whole status check if the number of failed execution does not exceed the failure threshold. Duration is available to both 'Synchronous' and 'Continuous' mode. A duration string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as '300ms', '-1.5h' or '2h45m'. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.",
						MarkdownDescription: "Duration defines the duration of the whole status check if the number of failed execution does not exceed the failure threshold. Duration is available to both 'Synchronous' and 'Continuous' mode. A duration string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as '300ms', '-1.5h' or '2h45m'. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"failure_threshold": schema.Int64Attribute{
						Description:         "FailureThreshold defines the minimum consecutive failure for the status check to be considered failed.",
						MarkdownDescription: "FailureThreshold defines the minimum consecutive failure for the status check to be considered failed.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"http": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"body": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"criteria": schema.SingleNestedAttribute{
								Description:         "Criteria defines how to determine the result of the status check.",
								MarkdownDescription: "Criteria defines how to determine the result of the status check.",
								Attributes: map[string]schema.Attribute{
									"status_code": schema.StringAttribute{
										Description:         "StatusCode defines the expected http status code for the request. A statusCode string could be a single code (e.g. 200), or an inclusive range (e.g. 200-400, both '200' and '400' are included).",
										MarkdownDescription: "StatusCode defines the expected http status code for the request. A statusCode string could be a single code (e.g. 200), or an inclusive range (e.g. 200-400, both '200' and '400' are included).",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"headers": schema.MapAttribute{
								Description:         "A Header represents the key-value pairs in an HTTP header.  The keys should be in canonical form, as returned by CanonicalHeaderKey.",
								MarkdownDescription: "A Header represents the key-value pairs in an HTTP header.  The keys should be in canonical form, as returned by CanonicalHeaderKey.",
								ElementType:         types.ListType{ElemType: types.StringType},
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"method": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"url": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"interval_seconds": schema.Int64Attribute{
						Description:         "IntervalSeconds defines how often (in seconds) to perform an execution of status check.",
						MarkdownDescription: "IntervalSeconds defines how often (in seconds) to perform an execution of status check.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"mode": schema.StringAttribute{
						Description:         "Mode defines the execution mode of the status check. Support type: Synchronous / Continuous",
						MarkdownDescription: "Mode defines the execution mode of the status check. Support type: Synchronous / Continuous",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"records_history_limit": schema.Int64Attribute{
						Description:         "RecordsHistoryLimit defines the number of record to retain.",
						MarkdownDescription: "RecordsHistoryLimit defines the number of record to retain.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"success_threshold": schema.Int64Attribute{
						Description:         "SuccessThreshold defines the minimum consecutive successes for the status check to be considered successful. SuccessThreshold only works for 'Synchronous' mode.",
						MarkdownDescription: "SuccessThreshold defines the minimum consecutive successes for the status check to be considered successful. SuccessThreshold only works for 'Synchronous' mode.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"timeout_seconds": schema.Int64Attribute{
						Description:         "TimeoutSeconds defines the number of seconds after which an execution of status check times out.",
						MarkdownDescription: "TimeoutSeconds defines the number of seconds after which an execution of status check times out.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"type": schema.StringAttribute{
						Description:         "Type defines the specific status check type. Support type: HTTP",
						MarkdownDescription: "Type defines the specific status check type. Support type: HTTP",
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

func (r *ChaosMeshOrgStatusCheckV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *ChaosMeshOrgStatusCheckV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_chaos_mesh_org_status_check_v1alpha1")

	var data ChaosMeshOrgStatusCheckV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "chaos-mesh.org", Version: "v1alpha1", Resource: "statuschecks"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		var statusError *k8sErrors.StatusError
		if errors.As(err, &statusError) {
			if statusError.Status().Code == http.StatusNotFound {
				response.Diagnostics.AddError(
					"Unable to find resource",
					fmt.Sprintf("The requested resource cannot be found. "+
						"Make sure that it does exist in your cluster and you have set the correct name and namespace configured.\n\n"+
						"Namespace: %s\n"+
						"Name: %s", data.Metadata.Namespace, data.Metadata.Name),
				)
				return
			}
		} else {
			response.Diagnostics.AddError(
				"Unable to GET resource",
				fmt.Sprintf("An unexpected error occurred while reading the resource. "+
					"Please report this issue to the provider developers.\n\n"+
					"GET Error (%T): %s", err, err.Error()),
			)
		}
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

	var readResponse ChaosMeshOrgStatusCheckV1Alpha1DataSourceData
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
	data.ApiVersion = pointer.String("chaos-mesh.org/v1alpha1")
	data.Kind = pointer.String("StatusCheck")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
