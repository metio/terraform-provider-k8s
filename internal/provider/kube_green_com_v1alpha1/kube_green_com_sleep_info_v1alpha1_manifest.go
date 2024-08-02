/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package kube_green_com_v1alpha1

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
	_ datasource.DataSource = &KubeGreenComSleepInfoV1Alpha1Manifest{}
)

func NewKubeGreenComSleepInfoV1Alpha1Manifest() datasource.DataSource {
	return &KubeGreenComSleepInfoV1Alpha1Manifest{}
}

type KubeGreenComSleepInfoV1Alpha1Manifest struct{}

type KubeGreenComSleepInfoV1Alpha1ManifestData struct {
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
		ExcludeRef *[]struct {
			ApiVersion  *string            `tfsdk:"api_version" json:"apiVersion,omitempty"`
			Kind        *string            `tfsdk:"kind" json:"kind,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			Name        *string            `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"exclude_ref" json:"excludeRef,omitempty"`
		IncludeRef *[]struct {
			ApiVersion  *string            `tfsdk:"api_version" json:"apiVersion,omitempty"`
			Kind        *string            `tfsdk:"kind" json:"kind,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			Name        *string            `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"include_ref" json:"includeRef,omitempty"`
		Patches *[]struct {
			Patch  *string `tfsdk:"patch" json:"patch,omitempty"`
			Target *struct {
				Group *string `tfsdk:"group" json:"group,omitempty"`
				Kind  *string `tfsdk:"kind" json:"kind,omitempty"`
			} `tfsdk:"target" json:"target,omitempty"`
		} `tfsdk:"patches" json:"patches,omitempty"`
		SleepAt             *string `tfsdk:"sleep_at" json:"sleepAt,omitempty"`
		SuspendCronJobs     *bool   `tfsdk:"suspend_cron_jobs" json:"suspendCronJobs,omitempty"`
		SuspendDeployments  *bool   `tfsdk:"suspend_deployments" json:"suspendDeployments,omitempty"`
		SuspendStatefulsets *bool   `tfsdk:"suspend_statefulsets" json:"suspendStatefulsets,omitempty"`
		TimeZone            *string `tfsdk:"time_zone" json:"timeZone,omitempty"`
		WakeUpAt            *string `tfsdk:"wake_up_at" json:"wakeUpAt,omitempty"`
		Weekdays            *string `tfsdk:"weekdays" json:"weekdays,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *KubeGreenComSleepInfoV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_kube_green_com_sleep_info_v1alpha1_manifest"
}

func (r *KubeGreenComSleepInfoV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "SleepInfo is the Schema for the sleepinfos API",
		MarkdownDescription: "SleepInfo is the Schema for the sleepinfos API",
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
				Description:         "SleepInfoSpec defines the desired state of SleepInfo",
				MarkdownDescription: "SleepInfoSpec defines the desired state of SleepInfo",
				Attributes: map[string]schema.Attribute{
					"exclude_ref": schema.ListNestedAttribute{
						Description:         "ExcludeRef define the resource to exclude from the sleep.",
						MarkdownDescription: "ExcludeRef define the resource to exclude from the sleep.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"api_version": schema.StringAttribute{
									Description:         "ApiVersion of the kubernetes resources.Supported api version is 'apps/v1'.",
									MarkdownDescription: "ApiVersion of the kubernetes resources.Supported api version is 'apps/v1'.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"kind": schema.StringAttribute{
									Description:         "Kind of the kubernetes resources of the specific version.Supported kind are 'Deployment' and 'CronJob'.",
									MarkdownDescription: "Kind of the kubernetes resources of the specific version.Supported kind are 'Deployment' and 'CronJob'.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"match_labels": schema.MapAttribute{
									Description:         "MatchLabels which identify the kubernetes resource by labels",
									MarkdownDescription: "MatchLabels which identify the kubernetes resource by labels",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "Name which identify the kubernetes resource.",
									MarkdownDescription: "Name which identify the kubernetes resource.",
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

					"include_ref": schema.ListNestedAttribute{
						Description:         "IncludeRef define the resource to include from the sleep.",
						MarkdownDescription: "IncludeRef define the resource to include from the sleep.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"api_version": schema.StringAttribute{
									Description:         "ApiVersion of the kubernetes resources.Supported api version is 'apps/v1'.",
									MarkdownDescription: "ApiVersion of the kubernetes resources.Supported api version is 'apps/v1'.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"kind": schema.StringAttribute{
									Description:         "Kind of the kubernetes resources of the specific version.Supported kind are 'Deployment' and 'CronJob'.",
									MarkdownDescription: "Kind of the kubernetes resources of the specific version.Supported kind are 'Deployment' and 'CronJob'.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"match_labels": schema.MapAttribute{
									Description:         "MatchLabels which identify the kubernetes resource by labels",
									MarkdownDescription: "MatchLabels which identify the kubernetes resource by labels",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "Name which identify the kubernetes resource.",
									MarkdownDescription: "Name which identify the kubernetes resource.",
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

					"patches": schema.ListNestedAttribute{
						Description:         "Patches is a list of json 6902 patches to apply to the target resources.",
						MarkdownDescription: "Patches is a list of json 6902 patches to apply to the target resources.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"patch": schema.StringAttribute{
									Description:         "Patch is the json6902 patch to apply to the target resource.",
									MarkdownDescription: "Patch is the json6902 patch to apply to the target resource.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"target": schema.SingleNestedAttribute{
									Description:         "Target is the target resource to patch.",
									MarkdownDescription: "Target is the target resource to patch.",
									Attributes: map[string]schema.Attribute{
										"group": schema.StringAttribute{
											Description:         "Group of the Kubernetes resources.",
											MarkdownDescription: "Group of the Kubernetes resources.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"kind": schema.StringAttribute{
											Description:         "Kind of the Kubernetes resources.",
											MarkdownDescription: "Kind of the Kubernetes resources.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},
									},
									Required: true,
									Optional: false,
									Computed: false,
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"sleep_at": schema.StringAttribute{
						Description:         "Hours:MinutesAccept cron schedule for both hour and minute.For example, *:*/2 is set to configure a run every even minute.",
						MarkdownDescription: "Hours:MinutesAccept cron schedule for both hour and minute.For example, *:*/2 is set to configure a run every even minute.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"suspend_cron_jobs": schema.BoolAttribute{
						Description:         "If SuspendCronjobs is set to true, on sleep the cronjobs of the namespace will be suspended.",
						MarkdownDescription: "If SuspendCronjobs is set to true, on sleep the cronjobs of the namespace will be suspended.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"suspend_deployments": schema.BoolAttribute{
						Description:         "If SuspendDeployments is set to false, on sleep the deployment of the namespace will not be suspended. By default Deployment will be suspended.",
						MarkdownDescription: "If SuspendDeployments is set to false, on sleep the deployment of the namespace will not be suspended. By default Deployment will be suspended.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"suspend_statefulsets": schema.BoolAttribute{
						Description:         "If SuspendStatefulSets is set to false, on sleep the statefulset of the namespace will not be suspended. By default StatefulSet will be suspended.",
						MarkdownDescription: "If SuspendStatefulSets is set to false, on sleep the statefulset of the namespace will not be suspended. By default StatefulSet will be suspended.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"time_zone": schema.StringAttribute{
						Description:         "Time zone to set the schedule, in IANA time zone identifier.It is not required, default to UTC.For example, for the Italy time zone set Europe/Rome.",
						MarkdownDescription: "Time zone to set the schedule, in IANA time zone identifier.It is not required, default to UTC.For example, for the Italy time zone set Europe/Rome.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"wake_up_at": schema.StringAttribute{
						Description:         "Hours:MinutesAccept cron schedule for both hour and minute.For example, *:*/2 is set to configure a run every even minute.It is not required.",
						MarkdownDescription: "Hours:MinutesAccept cron schedule for both hour and minute.For example, *:*/2 is set to configure a run every even minute.It is not required.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"weekdays": schema.StringAttribute{
						Description:         "Weekdays are in cron notation.For example, to configure a schedule from monday to friday, set it to '1-5'",
						MarkdownDescription: "Weekdays are in cron notation.For example, to configure a schedule from monday to friday, set it to '1-5'",
						Required:            true,
						Optional:            false,
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

func (r *KubeGreenComSleepInfoV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_kube_green_com_sleep_info_v1alpha1_manifest")

	var model KubeGreenComSleepInfoV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("kube-green.com/v1alpha1")
	model.Kind = pointer.String("SleepInfo")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
