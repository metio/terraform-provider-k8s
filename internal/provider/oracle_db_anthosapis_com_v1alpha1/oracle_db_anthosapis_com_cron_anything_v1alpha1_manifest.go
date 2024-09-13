/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package oracle_db_anthosapis_com_v1alpha1

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
	_ datasource.DataSource = &OracleDbAnthosapisComCronAnythingV1Alpha1Manifest{}
)

func NewOracleDbAnthosapisComCronAnythingV1Alpha1Manifest() datasource.DataSource {
	return &OracleDbAnthosapisComCronAnythingV1Alpha1Manifest{}
}

type OracleDbAnthosapisComCronAnythingV1Alpha1Manifest struct{}

type OracleDbAnthosapisComCronAnythingV1Alpha1ManifestData struct {
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
		CascadeDelete      *bool   `tfsdk:"cascade_delete" json:"cascadeDelete,omitempty"`
		ConcurrencyPolicy  *string `tfsdk:"concurrency_policy" json:"concurrencyPolicy,omitempty"`
		FinishableStrategy *struct {
			StringField *struct {
				FieldPath      *string   `tfsdk:"field_path" json:"fieldPath,omitempty"`
				FinishedValues *[]string `tfsdk:"finished_values" json:"finishedValues,omitempty"`
			} `tfsdk:"string_field" json:"stringField,omitempty"`
			TimestampField *struct {
				FieldPath *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
			} `tfsdk:"timestamp_field" json:"timestampField,omitempty"`
			Type *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"finishable_strategy" json:"finishableStrategy,omitempty"`
		ResourceBaseName        *string `tfsdk:"resource_base_name" json:"resourceBaseName,omitempty"`
		ResourceTimestampFormat *string `tfsdk:"resource_timestamp_format" json:"resourceTimestampFormat,omitempty"`
		Retention               *struct {
			HistoryCountLimit         *int64 `tfsdk:"history_count_limit" json:"historyCountLimit,omitempty"`
			HistoryTimeLimitSeconds   *int64 `tfsdk:"history_time_limit_seconds" json:"historyTimeLimitSeconds,omitempty"`
			ResourceTimestampStrategy *struct {
				Field *struct {
					FieldPath *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
				} `tfsdk:"field" json:"field,omitempty"`
				Type *string `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"resource_timestamp_strategy" json:"resourceTimestampStrategy,omitempty"`
		} `tfsdk:"retention" json:"retention,omitempty"`
		Schedule               *string            `tfsdk:"schedule" json:"schedule,omitempty"`
		Suspend                *bool              `tfsdk:"suspend" json:"suspend,omitempty"`
		Template               *map[string]string `tfsdk:"template" json:"template,omitempty"`
		TotalResourceLimit     *int64             `tfsdk:"total_resource_limit" json:"totalResourceLimit,omitempty"`
		TriggerDeadlineSeconds *int64             `tfsdk:"trigger_deadline_seconds" json:"triggerDeadlineSeconds,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *OracleDbAnthosapisComCronAnythingV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_oracle_db_anthosapis_com_cron_anything_v1alpha1_manifest"
}

func (r *OracleDbAnthosapisComCronAnythingV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "CronAnything is the Schema for the CronAnything API.",
		MarkdownDescription: "CronAnything is the Schema for the CronAnything API.",
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
				Description:         "CronAnythingSpec defines the desired state of CronAnything.",
				MarkdownDescription: "CronAnythingSpec defines the desired state of CronAnything.",
				Attributes: map[string]schema.Attribute{
					"cascade_delete": schema.BoolAttribute{
						Description:         "CascadeDelete tells CronAnything to set up owner references from the created resources to the CronAnything resource. This means that if the CronAnything resource is deleted, all resources created by it will also be deleted. This is an optional field that defaults to false.",
						MarkdownDescription: "CascadeDelete tells CronAnything to set up owner references from the created resources to the CronAnything resource. This means that if the CronAnything resource is deleted, all resources created by it will also be deleted. This is an optional field that defaults to false.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"concurrency_policy": schema.StringAttribute{
						Description:         "ConcurrencyPolicy specifies how to treat concurrent resources if the resource provides a status path that exposes completion. The default policy if not provided is to allow a new resource to be created even if an active resource already exists. If the resource doesn’t have an active/completed status, the only supported concurrency policy is to allow creating new resources. This field is mutable. If the policy is changed to a more stringent policy while multiple resources are active, it will not delete any existing resources. The exception is if a creation of a new resource is triggered and the policy has been changed to Replace. If multiple resources are active, they will all be deleted and replaced by a new resource.",
						MarkdownDescription: "ConcurrencyPolicy specifies how to treat concurrent resources if the resource provides a status path that exposes completion. The default policy if not provided is to allow a new resource to be created even if an active resource already exists. If the resource doesn’t have an active/completed status, the only supported concurrency policy is to allow creating new resources. This field is mutable. If the policy is changed to a more stringent policy while multiple resources are active, it will not delete any existing resources. The exception is if a creation of a new resource is triggered and the policy has been changed to Replace. If multiple resources are active, they will all be deleted and replaced by a new resource.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"finishable_strategy": schema.SingleNestedAttribute{
						Description:         "FinishableStrategy defines how the CronAnything controller an decide if a resource has completed. Some resources will do some work after they have been created and at some point be finished. Jobs are the most common example. If no strategy is defined, it is assumed that the resources never finish.",
						MarkdownDescription: "FinishableStrategy defines how the CronAnything controller an decide if a resource has completed. Some resources will do some work after they have been created and at some point be finished. Jobs are the most common example. If no strategy is defined, it is assumed that the resources never finish.",
						Attributes: map[string]schema.Attribute{
							"string_field": schema.SingleNestedAttribute{
								Description:         "StringField contains the details for how the CronAnything controller can find the string field on the resource needed to decide if the resource has completed. It also lists the values that mean the resource has completed.",
								MarkdownDescription: "StringField contains the details for how the CronAnything controller can find the string field on the resource needed to decide if the resource has completed. It also lists the values that mean the resource has completed.",
								Attributes: map[string]schema.Attribute{
									"field_path": schema.StringAttribute{
										Description:         "The path to the field on the resource that contains a string value.",
										MarkdownDescription: "The path to the field on the resource that contains a string value.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"finished_values": schema.ListAttribute{
										Description:         "The values of the field that means the resource has completed.",
										MarkdownDescription: "The values of the field that means the resource has completed.",
										ElementType:         types.StringType,
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"timestamp_field": schema.SingleNestedAttribute{
								Description:         "TimestampField contains the details for how the CronAnything controller can find the timestamp field on the resource in order to decide if the resource has completed.",
								MarkdownDescription: "TimestampField contains the details for how the CronAnything controller can find the timestamp field on the resource in order to decide if the resource has completed.",
								Attributes: map[string]schema.Attribute{
									"field_path": schema.StringAttribute{
										Description:         "The path to the field on the resource that contains the timestamp.",
										MarkdownDescription: "The path to the field on the resource that contains the timestamp.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"type": schema.StringAttribute{
								Description:         "Type tells which strategy should be used.",
								MarkdownDescription: "Type tells which strategy should be used.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"resource_base_name": schema.StringAttribute{
						Description:         "ResourceBaseName specifies the base name for the resources created by CronAnything, which will be named using the format <ResourceBaseName>-<Timestamp>. This field is optional, and the default is to use the name of the CronAnything resource as the ResourceBaseName.",
						MarkdownDescription: "ResourceBaseName specifies the base name for the resources created by CronAnything, which will be named using the format <ResourceBaseName>-<Timestamp>. This field is optional, and the default is to use the name of the CronAnything resource as the ResourceBaseName.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"resource_timestamp_format": schema.StringAttribute{
						Description:         "ResourceTimestampFormat defines the format of the timestamp in the name of Resources created by CronAnything <ResourceBaseName>-<Timestamp>. This field is optional, and the default is to format the timestamp as unix time. If provided, it must be compatible with time.Format in golang.",
						MarkdownDescription: "ResourceTimestampFormat defines the format of the timestamp in the name of Resources created by CronAnything <ResourceBaseName>-<Timestamp>. This field is optional, and the default is to format the timestamp as unix time. If provided, it must be compatible with time.Format in golang.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"retention": schema.SingleNestedAttribute{
						Description:         "Retention defines the retention policy for resources created by CronAnything. If no retention policy is defined, CronAnything will never delete resources, so cleanup must be handled through some other process.",
						MarkdownDescription: "Retention defines the retention policy for resources created by CronAnything. If no retention policy is defined, CronAnything will never delete resources, so cleanup must be handled through some other process.",
						Attributes: map[string]schema.Attribute{
							"history_count_limit": schema.Int64Attribute{
								Description:         "The number of completed resources to keep before deleting them. This only affects finishable resources and the default value is 3. This field is mutable and if it is changed to a number lower than the current number of finished resources, the oldest ones will eventually be deleted until the number of finished resources matches the limit.",
								MarkdownDescription: "The number of completed resources to keep before deleting them. This only affects finishable resources and the default value is 3. This field is mutable and if it is changed to a number lower than the current number of finished resources, the oldest ones will eventually be deleted until the number of finished resources matches the limit.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"history_time_limit_seconds": schema.Int64Attribute{
								Description:         "The time since completion that a resource is kept before deletion. This only affects finishable resources. This does not have any default value and if it is not provided, HistoryCountLimit will be used to prune completed resources. If both HistoryCountLimit and HistoryTimeLimitSeconds are set, it is treated as an OR operation.",
								MarkdownDescription: "The time since completion that a resource is kept before deletion. This only affects finishable resources. This does not have any default value and if it is not provided, HistoryCountLimit will be used to prune completed resources. If both HistoryCountLimit and HistoryTimeLimitSeconds are set, it is treated as an OR operation.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"resource_timestamp_strategy": schema.SingleNestedAttribute{
								Description:         "ResourceTimestampStrategy specifies how the CronAnything controller can find the age of a resource. This is needed to support retention.",
								MarkdownDescription: "ResourceTimestampStrategy specifies how the CronAnything controller can find the age of a resource. This is needed to support retention.",
								Attributes: map[string]schema.Attribute{
									"field": schema.SingleNestedAttribute{
										Description:         "FieldResourceTimestampStrategy specifies how the CronAnything controller can find the timestamp for the resource from a field.",
										MarkdownDescription: "FieldResourceTimestampStrategy specifies how the CronAnything controller can find the timestamp for the resource from a field.",
										Attributes: map[string]schema.Attribute{
											"field_path": schema.StringAttribute{
												Description:         "The path to the field on the resource that contains the timestamp.",
												MarkdownDescription: "The path to the field on the resource that contains the timestamp.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"type": schema.StringAttribute{
										Description:         "Type tells which strategy should be used.",
										MarkdownDescription: "Type tells which strategy should be used.",
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
						Required: false,
						Optional: true,
						Computed: false,
					},

					"schedule": schema.StringAttribute{
						Description:         "Schedule defines a time-based schedule, e.g., a standard cron schedule such as “@every 10m”. This field is mandatory and mutable. If it is changed, resources will simply be created at the new interval from then on.",
						MarkdownDescription: "Schedule defines a time-based schedule, e.g., a standard cron schedule such as “@every 10m”. This field is mandatory and mutable. If it is changed, resources will simply be created at the new interval from then on.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"suspend": schema.BoolAttribute{
						Description:         "Suspend tells the controller to suspend creation of additional resources. The default value is false. This field is mutable. It will not affect any existing resources, but only affect creation of additional resources.",
						MarkdownDescription: "Suspend tells the controller to suspend creation of additional resources. The default value is false. This field is mutable. It will not affect any existing resources, but only affect creation of additional resources.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"template": schema.MapAttribute{
						Description:         "Template is a template of a resource type for which instances are to be created on the given schedule. This field is mandatory and it must contain a valid template for an existing apiVersion and kind in the cluster. It is immutable, so if the template needs to change, the whole CronAnything resource should be replaced.",
						MarkdownDescription: "Template is a template of a resource type for which instances are to be created on the given schedule. This field is mandatory and it must contain a valid template for an existing apiVersion and kind in the cluster. It is immutable, so if the template needs to change, the whole CronAnything resource should be replaced.",
						ElementType:         types.StringType,
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"total_resource_limit": schema.Int64Attribute{
						Description:         "TotalResourceLimit specifies the total number of children allowed for a particular CronAnything resource. If this limit is reached, no additional resources will be created. This limit is mostly meant to avoid runaway creation of resources that could bring down the cluster. Both finished and unfinished resources count against this limit. This field is mutable. If it is changed to a lower value than the existing number of resources, none of the existing resources will be deleted as a result, but no additional resources will be created until the number of child resources goes below the limit. The field is optional with a default value of 100.",
						MarkdownDescription: "TotalResourceLimit specifies the total number of children allowed for a particular CronAnything resource. If this limit is reached, no additional resources will be created. This limit is mostly meant to avoid runaway creation of resources that could bring down the cluster. Both finished and unfinished resources count against this limit. This field is mutable. If it is changed to a lower value than the existing number of resources, none of the existing resources will be deleted as a result, but no additional resources will be created until the number of child resources goes below the limit. The field is optional with a default value of 100.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"trigger_deadline_seconds": schema.Int64Attribute{
						Description:         "TriggerDeadlineSeconds defines Deadline in seconds for creating the resource if it missed the scheduled time. If no deadline is provided, the resource will be created no matter how far after the scheduled time. If multiple triggers were missed, only the last will be triggered and only one resource will be created. This field is mutable and changing it will affect the creation of new resources from that point in time.",
						MarkdownDescription: "TriggerDeadlineSeconds defines Deadline in seconds for creating the resource if it missed the scheduled time. If no deadline is provided, the resource will be created no matter how far after the scheduled time. If multiple triggers were missed, only the last will be triggered and only one resource will be created. This field is mutable and changing it will affect the creation of new resources from that point in time.",
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

func (r *OracleDbAnthosapisComCronAnythingV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_oracle_db_anthosapis_com_cron_anything_v1alpha1_manifest")

	var model OracleDbAnthosapisComCronAnythingV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("oracle.db.anthosapis.com/v1alpha1")
	model.Kind = pointer.String("CronAnything")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
