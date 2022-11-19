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

type SchedulingKoordinatorShPodMigrationJobV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*SchedulingKoordinatorShPodMigrationJobV1Alpha1Resource)(nil)
)

type SchedulingKoordinatorShPodMigrationJobV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type SchedulingKoordinatorShPodMigrationJobV1Alpha1GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		DeleteOptions *struct {
			ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion,omitempty"`

			DryRun *[]string `tfsdk:"dry_run" yaml:"dryRun,omitempty"`

			GracePeriodSeconds *int64 `tfsdk:"grace_period_seconds" yaml:"gracePeriodSeconds,omitempty"`

			Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

			OrphanDependents *bool `tfsdk:"orphan_dependents" yaml:"orphanDependents,omitempty"`

			Preconditions *struct {
				ResourceVersion *string `tfsdk:"resource_version" yaml:"resourceVersion,omitempty"`

				Uid *string `tfsdk:"uid" yaml:"uid,omitempty"`
			} `tfsdk:"preconditions" yaml:"preconditions,omitempty"`

			PropagationPolicy *string `tfsdk:"propagation_policy" yaml:"propagationPolicy,omitempty"`
		} `tfsdk:"delete_options" yaml:"deleteOptions,omitempty"`

		Mode *string `tfsdk:"mode" yaml:"mode,omitempty"`

		Paused *bool `tfsdk:"paused" yaml:"paused,omitempty"`

		PodRef *struct {
			ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion,omitempty"`

			FieldPath *string `tfsdk:"field_path" yaml:"fieldPath,omitempty"`

			Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

			ResourceVersion *string `tfsdk:"resource_version" yaml:"resourceVersion,omitempty"`

			Uid *string `tfsdk:"uid" yaml:"uid,omitempty"`
		} `tfsdk:"pod_ref" yaml:"podRef,omitempty"`

		ReservationOptions *struct {
			PreemptionOptions *map[string]string `tfsdk:"preemption_options" yaml:"preemptionOptions,omitempty"`

			ReservationRef *struct {
				ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion,omitempty"`

				FieldPath *string `tfsdk:"field_path" yaml:"fieldPath,omitempty"`

				Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

				ResourceVersion *string `tfsdk:"resource_version" yaml:"resourceVersion,omitempty"`

				Uid *string `tfsdk:"uid" yaml:"uid,omitempty"`
			} `tfsdk:"reservation_ref" yaml:"reservationRef,omitempty"`

			Template utilities.Dynamic `tfsdk:"template" yaml:"template,omitempty"`
		} `tfsdk:"reservation_options" yaml:"reservationOptions,omitempty"`

		Ttl *string `tfsdk:"ttl" yaml:"ttl,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewSchedulingKoordinatorShPodMigrationJobV1Alpha1Resource() resource.Resource {
	return &SchedulingKoordinatorShPodMigrationJobV1Alpha1Resource{}
}

func (r *SchedulingKoordinatorShPodMigrationJobV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_scheduling_koordinator_sh_pod_migration_job_v1alpha1"
}

func (r *SchedulingKoordinatorShPodMigrationJobV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "",
		MarkdownDescription: "",
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
				Description:         "",
				MarkdownDescription: "",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"delete_options": {
						Description:         "DeleteOptions defines the deleting options for the migrated Pod and preempted Pods",
						MarkdownDescription: "DeleteOptions defines the deleting options for the migrated Pod and preempted Pods",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"api_version": {
								Description:         "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
								MarkdownDescription: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"dry_run": {
								Description:         "When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed",
								MarkdownDescription: "When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"grace_period_seconds": {
								Description:         "The duration in seconds before the object should be deleted. Value must be non-negative integer. The value zero indicates delete immediately. If this value is nil, the default grace period for the specified type will be used. Defaults to a per object value if not specified. zero means delete immediately.",
								MarkdownDescription: "The duration in seconds before the object should be deleted. Value must be non-negative integer. The value zero indicates delete immediately. If this value is nil, the default grace period for the specified type will be used. Defaults to a per object value if not specified. zero means delete immediately.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"kind": {
								Description:         "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
								MarkdownDescription: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"orphan_dependents": {
								Description:         "Deprecated: please use the PropagationPolicy, this field will be deprecated in 1.7. Should the dependent objects be orphaned. If true/false, the 'orphan' finalizer will be added to/removed from the object's finalizers list. Either this field or PropagationPolicy may be set, but not both.",
								MarkdownDescription: "Deprecated: please use the PropagationPolicy, this field will be deprecated in 1.7. Should the dependent objects be orphaned. If true/false, the 'orphan' finalizer will be added to/removed from the object's finalizers list. Either this field or PropagationPolicy may be set, but not both.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"preconditions": {
								Description:         "Must be fulfilled before a deletion is carried out. If not possible, a 409 Conflict status will be returned.",
								MarkdownDescription: "Must be fulfilled before a deletion is carried out. If not possible, a 409 Conflict status will be returned.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"resource_version": {
										Description:         "Specifies the target ResourceVersion",
										MarkdownDescription: "Specifies the target ResourceVersion",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"uid": {
										Description:         "Specifies the target UID.",
										MarkdownDescription: "Specifies the target UID.",

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

							"propagation_policy": {
								Description:         "Whether and how garbage collection will be performed. Either this field or OrphanDependents may be set, but not both. The default policy is decided by the existing finalizer set in the metadata.finalizers and the resource-specific default policy. Acceptable values are: 'Orphan' - orphan the dependents; 'Background' - allow the garbage collector to delete the dependents in the background; 'Foreground' - a cascading policy that deletes all dependents in the foreground.",
								MarkdownDescription: "Whether and how garbage collection will be performed. Either this field or OrphanDependents may be set, but not both. The default policy is decided by the existing finalizer set in the metadata.finalizers and the resource-specific default policy. Acceptable values are: 'Orphan' - orphan the dependents; 'Background' - allow the garbage collector to delete the dependents in the background; 'Foreground' - a cascading policy that deletes all dependents in the foreground.",

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

					"mode": {
						Description:         "Mode represents the operating mode of the Job Default is PodMigrationJobModeReservationFirst",
						MarkdownDescription: "Mode represents the operating mode of the Job Default is PodMigrationJobModeReservationFirst",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"paused": {
						Description:         "Paused indicates whether the PodMigrationJob should to work or not. Default is false",
						MarkdownDescription: "Paused indicates whether the PodMigrationJob should to work or not. Default is false",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"pod_ref": {
						Description:         "PodRef represents the Pod that be migrated",
						MarkdownDescription: "PodRef represents the Pod that be migrated",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"api_version": {
								Description:         "API version of the referent.",
								MarkdownDescription: "API version of the referent.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"field_path": {
								Description:         "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object. TODO: this design is not final and this field is subject to change in the future.",
								MarkdownDescription: "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object. TODO: this design is not final and this field is subject to change in the future.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"kind": {
								Description:         "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
								MarkdownDescription: "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"name": {
								Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
								MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"namespace": {
								Description:         "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
								MarkdownDescription: "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"resource_version": {
								Description:         "Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
								MarkdownDescription: "Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"uid": {
								Description:         "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
								MarkdownDescription: "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: true,
						Optional: false,
						Computed: false,
					},

					"reservation_options": {
						Description:         "ReservationOptions defines the Reservation options for migrated Pod",
						MarkdownDescription: "ReservationOptions defines the Reservation options for migrated Pod",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"preemption_options": {
								Description:         "PreemptionOption decides whether to preempt other Pods. The preemption is safe and reserves resources for preempted Pods.",
								MarkdownDescription: "PreemptionOption decides whether to preempt other Pods. The preemption is safe and reserves resources for preempted Pods.",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"reservation_ref": {
								Description:         "ReservationRef if specified, PodMigrationJob will check if the status of Reservation is available. ReservationRef if not specified, PodMigrationJob controller will create Reservation by Template, and update the ReservationRef to reference the Reservation",
								MarkdownDescription: "ReservationRef if specified, PodMigrationJob will check if the status of Reservation is available. ReservationRef if not specified, PodMigrationJob controller will create Reservation by Template, and update the ReservationRef to reference the Reservation",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"api_version": {
										Description:         "API version of the referent.",
										MarkdownDescription: "API version of the referent.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"field_path": {
										Description:         "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object. TODO: this design is not final and this field is subject to change in the future.",
										MarkdownDescription: "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object. TODO: this design is not final and this field is subject to change in the future.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"kind": {
										Description:         "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
										MarkdownDescription: "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"name": {
										Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
										MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"namespace": {
										Description:         "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
										MarkdownDescription: "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"resource_version": {
										Description:         "Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
										MarkdownDescription: "Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"uid": {
										Description:         "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
										MarkdownDescription: "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",

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

							"template": {
								Description:         "Template is the object that describes the Reservation that will be created if not specified ReservationRef",
								MarkdownDescription: "Template is the object that describes the Reservation that will be created if not specified ReservationRef",

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

					"ttl": {
						Description:         "TTL controls the PodMigrationJob timeout duration.",
						MarkdownDescription: "TTL controls the PodMigrationJob timeout duration.",

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

func (r *SchedulingKoordinatorShPodMigrationJobV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_scheduling_koordinator_sh_pod_migration_job_v1alpha1")

	var state SchedulingKoordinatorShPodMigrationJobV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel SchedulingKoordinatorShPodMigrationJobV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("scheduling.koordinator.sh/v1alpha1")
	goModel.Kind = utilities.Ptr("PodMigrationJob")

	state.Id = types.Int64Value(time.Now().UnixNano())
	state.ApiVersion = types.StringValue(*goModel.ApiVersion)
	state.Kind = types.StringValue(*goModel.Kind)

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.StringValue(string(marshal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *SchedulingKoordinatorShPodMigrationJobV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_scheduling_koordinator_sh_pod_migration_job_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *SchedulingKoordinatorShPodMigrationJobV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_scheduling_koordinator_sh_pod_migration_job_v1alpha1")

	var state SchedulingKoordinatorShPodMigrationJobV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel SchedulingKoordinatorShPodMigrationJobV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("scheduling.koordinator.sh/v1alpha1")
	goModel.Kind = utilities.Ptr("PodMigrationJob")

	state.Id = types.Int64Value(time.Now().UnixNano())
	state.ApiVersion = types.StringValue(*goModel.ApiVersion)
	state.Kind = types.StringValue(*goModel.Kind)

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.StringValue(string(marshal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *SchedulingKoordinatorShPodMigrationJobV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_scheduling_koordinator_sh_pod_migration_job_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
