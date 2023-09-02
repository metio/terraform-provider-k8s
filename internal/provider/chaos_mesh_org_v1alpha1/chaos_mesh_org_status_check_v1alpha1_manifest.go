/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package chaos_mesh_org_v1alpha1

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
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
	_ datasource.DataSource = &ChaosMeshOrgStatusCheckV1Alpha1Manifest{}
)

func NewChaosMeshOrgStatusCheckV1Alpha1Manifest() datasource.DataSource {
	return &ChaosMeshOrgStatusCheckV1Alpha1Manifest{}
}

type ChaosMeshOrgStatusCheckV1Alpha1Manifest struct{}

type ChaosMeshOrgStatusCheckV1Alpha1ManifestData struct {
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

func (r *ChaosMeshOrgStatusCheckV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_chaos_mesh_org_status_check_v1alpha1_manifest"
}

func (r *ChaosMeshOrgStatusCheckV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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
				Description:         "Spec defines the behavior of a status check",
				MarkdownDescription: "Spec defines the behavior of a status check",
				Attributes: map[string]schema.Attribute{
					"duration": schema.StringAttribute{
						Description:         "Duration defines the duration of the whole status check if the number of failed execution does not exceed the failure threshold. Duration is available to both 'Synchronous' and 'Continuous' mode. A duration string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as '300ms', '-1.5h' or '2h45m'. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.",
						MarkdownDescription: "Duration defines the duration of the whole status check if the number of failed execution does not exceed the failure threshold. Duration is available to both 'Synchronous' and 'Continuous' mode. A duration string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as '300ms', '-1.5h' or '2h45m'. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"failure_threshold": schema.Int64Attribute{
						Description:         "FailureThreshold defines the minimum consecutive failure for the status check to be considered failed.",
						MarkdownDescription: "FailureThreshold defines the minimum consecutive failure for the status check to be considered failed.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(1),
						},
					},

					"http": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"body": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"criteria": schema.SingleNestedAttribute{
								Description:         "Criteria defines how to determine the result of the status check.",
								MarkdownDescription: "Criteria defines how to determine the result of the status check.",
								Attributes: map[string]schema.Attribute{
									"status_code": schema.StringAttribute{
										Description:         "StatusCode defines the expected http status code for the request. A statusCode string could be a single code (e.g. 200), or an inclusive range (e.g. 200-400, both '200' and '400' are included).",
										MarkdownDescription: "StatusCode defines the expected http status code for the request. A statusCode string could be a single code (e.g. 200), or an inclusive range (e.g. 200-400, both '200' and '400' are included).",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: true,
								Optional: false,
								Computed: false,
							},

							"headers": schema.MapAttribute{
								Description:         "A Header represents the key-value pairs in an HTTP header.  The keys should be in canonical form, as returned by CanonicalHeaderKey.",
								MarkdownDescription: "A Header represents the key-value pairs in an HTTP header.  The keys should be in canonical form, as returned by CanonicalHeaderKey.",
								ElementType:         types.ListType{ElemType: types.StringType},
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"method": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("GET", "POST"),
								},
							},

							"url": schema.StringAttribute{
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

					"interval_seconds": schema.Int64Attribute{
						Description:         "IntervalSeconds defines how often (in seconds) to perform an execution of status check.",
						MarkdownDescription: "IntervalSeconds defines how often (in seconds) to perform an execution of status check.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(1),
						},
					},

					"mode": schema.StringAttribute{
						Description:         "Mode defines the execution mode of the status check. Support type: Synchronous / Continuous",
						MarkdownDescription: "Mode defines the execution mode of the status check. Support type: Synchronous / Continuous",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("Synchronous", "Continuous"),
						},
					},

					"records_history_limit": schema.Int64Attribute{
						Description:         "RecordsHistoryLimit defines the number of record to retain.",
						MarkdownDescription: "RecordsHistoryLimit defines the number of record to retain.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(1),
							int64validator.AtMost(1000),
						},
					},

					"success_threshold": schema.Int64Attribute{
						Description:         "SuccessThreshold defines the minimum consecutive successes for the status check to be considered successful. SuccessThreshold only works for 'Synchronous' mode.",
						MarkdownDescription: "SuccessThreshold defines the minimum consecutive successes for the status check to be considered successful. SuccessThreshold only works for 'Synchronous' mode.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(1),
						},
					},

					"timeout_seconds": schema.Int64Attribute{
						Description:         "TimeoutSeconds defines the number of seconds after which an execution of status check times out.",
						MarkdownDescription: "TimeoutSeconds defines the number of seconds after which an execution of status check times out.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(1),
						},
					},

					"type": schema.StringAttribute{
						Description:         "Type defines the specific status check type. Support type: HTTP",
						MarkdownDescription: "Type defines the specific status check type. Support type: HTTP",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("HTTP"),
						},
					},
				},
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *ChaosMeshOrgStatusCheckV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_chaos_mesh_org_status_check_v1alpha1_manifest")

	var model ChaosMeshOrgStatusCheckV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Name, model.Metadata.Namespace))
	model.ApiVersion = pointer.String("chaos-mesh.org/v1alpha1")
	model.Kind = pointer.String("StatusCheck")

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
