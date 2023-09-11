/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package autoscaling_karmada_io_v1alpha1

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
	_ datasource.DataSource              = &AutoscalingKarmadaIoCronFederatedHpaV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &AutoscalingKarmadaIoCronFederatedHpaV1Alpha1DataSource{}
)

func NewAutoscalingKarmadaIoCronFederatedHpaV1Alpha1DataSource() datasource.DataSource {
	return &AutoscalingKarmadaIoCronFederatedHpaV1Alpha1DataSource{}
}

type AutoscalingKarmadaIoCronFederatedHpaV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type AutoscalingKarmadaIoCronFederatedHpaV1Alpha1DataSourceData struct {
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
		Rules *[]struct {
			FailedHistoryLimit     *int64  `tfsdk:"failed_history_limit" json:"failedHistoryLimit,omitempty"`
			Name                   *string `tfsdk:"name" json:"name,omitempty"`
			Schedule               *string `tfsdk:"schedule" json:"schedule,omitempty"`
			SuccessfulHistoryLimit *int64  `tfsdk:"successful_history_limit" json:"successfulHistoryLimit,omitempty"`
			Suspend                *bool   `tfsdk:"suspend" json:"suspend,omitempty"`
			TargetMaxReplicas      *int64  `tfsdk:"target_max_replicas" json:"targetMaxReplicas,omitempty"`
			TargetMinReplicas      *int64  `tfsdk:"target_min_replicas" json:"targetMinReplicas,omitempty"`
			TargetReplicas         *int64  `tfsdk:"target_replicas" json:"targetReplicas,omitempty"`
			TimeZone               *string `tfsdk:"time_zone" json:"timeZone,omitempty"`
		} `tfsdk:"rules" json:"rules,omitempty"`
		ScaleTargetRef *struct {
			ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
			Kind       *string `tfsdk:"kind" json:"kind,omitempty"`
			Name       *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"scale_target_ref" json:"scaleTargetRef,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AutoscalingKarmadaIoCronFederatedHpaV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_autoscaling_karmada_io_cron_federated_hpa_v1alpha1"
}

func (r *AutoscalingKarmadaIoCronFederatedHpaV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "CronFederatedHPA represents a collection of repeating schedule to scale replica number of a specific workload. It can scale any resource implementing the scale subresource as well as FederatedHPA.",
		MarkdownDescription: "CronFederatedHPA represents a collection of repeating schedule to scale replica number of a specific workload. It can scale any resource implementing the scale subresource as well as FederatedHPA.",
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
				Description:         "Spec is the specification of the CronFederatedHPA.",
				MarkdownDescription: "Spec is the specification of the CronFederatedHPA.",
				Attributes: map[string]schema.Attribute{
					"rules": schema.ListNestedAttribute{
						Description:         "Rules contains a collection of schedules that declares when and how the referencing target resource should be scaled.",
						MarkdownDescription: "Rules contains a collection of schedules that declares when and how the referencing target resource should be scaled.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"failed_history_limit": schema.Int64Attribute{
									Description:         "FailedHistoryLimit represents the count of failed execution items for each rule. The value must be a positive integer. It defaults to 3.",
									MarkdownDescription: "FailedHistoryLimit represents the count of failed execution items for each rule. The value must be a positive integer. It defaults to 3.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"name": schema.StringAttribute{
									Description:         "Name of the rule. Each rule in a CronFederatedHPA must have a unique name.  Note: the name will be used as an identifier to record its execution history. Changing the name will be considered as deleting the old rule and adding a new rule, that means the original execution history will be discarded.",
									MarkdownDescription: "Name of the rule. Each rule in a CronFederatedHPA must have a unique name.  Note: the name will be used as an identifier to record its execution history. Changing the name will be considered as deleting the old rule and adding a new rule, that means the original execution history will be discarded.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"schedule": schema.StringAttribute{
									Description:         "Schedule is the cron expression that represents a periodical time. The syntax follows https://kubernetes.io/docs/concepts/workloads/controllers/cron-jobs/#schedule-syntax.",
									MarkdownDescription: "Schedule is the cron expression that represents a periodical time. The syntax follows https://kubernetes.io/docs/concepts/workloads/controllers/cron-jobs/#schedule-syntax.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"successful_history_limit": schema.Int64Attribute{
									Description:         "SuccessfulHistoryLimit represents the count of successful execution items for each rule. The value must be a positive integer. It defaults to 3.",
									MarkdownDescription: "SuccessfulHistoryLimit represents the count of successful execution items for each rule. The value must be a positive integer. It defaults to 3.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"suspend": schema.BoolAttribute{
									Description:         "Suspend tells the controller to suspend subsequent executions. Defaults to false.",
									MarkdownDescription: "Suspend tells the controller to suspend subsequent executions. Defaults to false.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"target_max_replicas": schema.Int64Attribute{
									Description:         "TargetMaxReplicas is the target MaxReplicas to be set for FederatedHPA. Only needed when referencing resource is FederatedHPA. TargetMinReplicas and TargetMaxReplicas can be specified together or either one can be specified alone. nil means the MaxReplicas(.spec.maxReplicas) of the referencing FederatedHPA will not be updated.",
									MarkdownDescription: "TargetMaxReplicas is the target MaxReplicas to be set for FederatedHPA. Only needed when referencing resource is FederatedHPA. TargetMinReplicas and TargetMaxReplicas can be specified together or either one can be specified alone. nil means the MaxReplicas(.spec.maxReplicas) of the referencing FederatedHPA will not be updated.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"target_min_replicas": schema.Int64Attribute{
									Description:         "TargetMinReplicas is the target MinReplicas to be set for FederatedHPA. Only needed when referencing resource is FederatedHPA. TargetMinReplicas and TargetMaxReplicas can be specified together or either one can be specified alone. nil means the MinReplicas(.spec.minReplicas) of the referencing FederatedHPA will not be updated.",
									MarkdownDescription: "TargetMinReplicas is the target MinReplicas to be set for FederatedHPA. Only needed when referencing resource is FederatedHPA. TargetMinReplicas and TargetMaxReplicas can be specified together or either one can be specified alone. nil means the MinReplicas(.spec.minReplicas) of the referencing FederatedHPA will not be updated.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"target_replicas": schema.Int64Attribute{
									Description:         "TargetReplicas is the target replicas to be scaled for resources referencing by ScaleTargetRef of this CronFederatedHPA. Only needed when referencing resource is not FederatedHPA.",
									MarkdownDescription: "TargetReplicas is the target replicas to be scaled for resources referencing by ScaleTargetRef of this CronFederatedHPA. Only needed when referencing resource is not FederatedHPA.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"time_zone": schema.StringAttribute{
									Description:         "TimeZone for the giving schedule. If not specified, this will default to the time zone of the karmada-controller-manager process. Invalid TimeZone will be rejected when applying by karmada-webhook. see https://en.wikipedia.org/wiki/List_of_tz_database_time_zones for the all timezones.",
									MarkdownDescription: "TimeZone for the giving schedule. If not specified, this will default to the time zone of the karmada-controller-manager process. Invalid TimeZone will be rejected when applying by karmada-webhook. see https://en.wikipedia.org/wiki/List_of_tz_database_time_zones for the all timezones.",
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

					"scale_target_ref": schema.SingleNestedAttribute{
						Description:         "ScaleTargetRef points to the target resource to scale. Target resource could be any resource that implementing the scale subresource like Deployment, or FederatedHPA.",
						MarkdownDescription: "ScaleTargetRef points to the target resource to scale. Target resource could be any resource that implementing the scale subresource like Deployment, or FederatedHPA.",
						Attributes: map[string]schema.Attribute{
							"api_version": schema.StringAttribute{
								Description:         "apiVersion is the API version of the referent",
								MarkdownDescription: "apiVersion is the API version of the referent",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"kind": schema.StringAttribute{
								Description:         "kind is the kind of the referent; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
								MarkdownDescription: "kind is the kind of the referent; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"name": schema.StringAttribute{
								Description:         "name is the name of the referent; More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
								MarkdownDescription: "name is the name of the referent; More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
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
				Required: false,
				Optional: false,
				Computed: true,
			},
		},
	}
}

func (r *AutoscalingKarmadaIoCronFederatedHpaV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *AutoscalingKarmadaIoCronFederatedHpaV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_autoscaling_karmada_io_cron_federated_hpa_v1alpha1")

	var data AutoscalingKarmadaIoCronFederatedHpaV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "autoscaling.karmada.io", Version: "v1alpha1", Resource: "cronfederatedhpas"}).
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

	var readResponse AutoscalingKarmadaIoCronFederatedHpaV1Alpha1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Namespace, data.Metadata.Name))
	data.ApiVersion = pointer.String("autoscaling.karmada.io/v1alpha1")
	data.Kind = pointer.String("CronFederatedHPA")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
