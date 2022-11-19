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

type EventsK8SIoEventV1Resource struct{}

var (
	_ resource.Resource = (*EventsK8SIoEventV1Resource)(nil)
)

type EventsK8SIoEventV1TerraformModel struct {
	Id                       types.Int64  `tfsdk:"id"`
	YAML                     types.String `tfsdk:"yaml"`
	ApiVersion               types.String `tfsdk:"api_version"`
	Kind                     types.String `tfsdk:"kind"`
	Metadata                 types.Object `tfsdk:"metadata"`
	Action                   types.String `tfsdk:"action"`
	DeprecatedCount          types.Int64  `tfsdk:"deprecated_count"`
	DeprecatedFirstTimestamp types.String `tfsdk:"deprecated_first_timestamp"`
	DeprecatedLastTimestamp  types.String `tfsdk:"deprecated_last_timestamp"`
	DeprecatedSource         types.Object `tfsdk:"deprecated_source"`
	EventTime                types.String `tfsdk:"event_time"`
	Note                     types.String `tfsdk:"note"`
	Reason                   types.String `tfsdk:"reason"`
	Regarding                types.Object `tfsdk:"regarding"`
	Related                  types.Object `tfsdk:"related"`
	ReportingController      types.String `tfsdk:"reporting_controller"`
	ReportingInstance        types.String `tfsdk:"reporting_instance"`
	Series                   types.Object `tfsdk:"series"`
	Type                     types.String `tfsdk:"type"`
}

type EventsK8SIoEventV1GoModel struct {
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

	Action *string `tfsdk:"action" yaml:"action,omitempty"`

	DeprecatedCount *int64 `tfsdk:"deprecated_count" yaml:"deprecatedCount,omitempty"`

	DeprecatedFirstTimestamp *string `tfsdk:"deprecated_first_timestamp" yaml:"deprecatedFirstTimestamp,omitempty"`

	DeprecatedLastTimestamp *string `tfsdk:"deprecated_last_timestamp" yaml:"deprecatedLastTimestamp,omitempty"`

	DeprecatedSource *struct {
		Component *string `tfsdk:"component" yaml:"component,omitempty"`

		Host *string `tfsdk:"host" yaml:"host,omitempty"`
	} `tfsdk:"deprecated_source" yaml:"deprecatedSource,omitempty"`

	EventTime *string `tfsdk:"event_time" yaml:"eventTime,omitempty"`

	Note *string `tfsdk:"note" yaml:"note,omitempty"`

	Reason *string `tfsdk:"reason" yaml:"reason,omitempty"`

	Regarding *struct {
		ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion,omitempty"`

		FieldPath *string `tfsdk:"field_path" yaml:"fieldPath,omitempty"`

		Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

		Name *string `tfsdk:"name" yaml:"name,omitempty"`

		Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

		ResourceVersion *string `tfsdk:"resource_version" yaml:"resourceVersion,omitempty"`

		Uid *string `tfsdk:"uid" yaml:"uid,omitempty"`
	} `tfsdk:"regarding" yaml:"regarding,omitempty"`

	Related *struct {
		ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion,omitempty"`

		FieldPath *string `tfsdk:"field_path" yaml:"fieldPath,omitempty"`

		Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

		Name *string `tfsdk:"name" yaml:"name,omitempty"`

		Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

		ResourceVersion *string `tfsdk:"resource_version" yaml:"resourceVersion,omitempty"`

		Uid *string `tfsdk:"uid" yaml:"uid,omitempty"`
	} `tfsdk:"related" yaml:"related,omitempty"`

	ReportingController *string `tfsdk:"reporting_controller" yaml:"reportingController,omitempty"`

	ReportingInstance *string `tfsdk:"reporting_instance" yaml:"reportingInstance,omitempty"`

	Series *struct {
		Count *int64 `tfsdk:"count" yaml:"count,omitempty"`

		LastObservedTime *string `tfsdk:"last_observed_time" yaml:"lastObservedTime,omitempty"`
	} `tfsdk:"series" yaml:"series,omitempty"`

	Type *string `tfsdk:"type" yaml:"type,omitempty"`
}

func NewEventsK8SIoEventV1Resource() resource.Resource {
	return &EventsK8SIoEventV1Resource{}
}

func (r *EventsK8SIoEventV1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_events_k8s_io_event_v1"
}

func (r *EventsK8SIoEventV1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "Event is a report of an event somewhere in the cluster. It generally denotes some state change in the system. Events have a limited retention time and triggers and messages may evolve with time.  Event consumers should not rely on the timing of an event with a given Reason reflecting a consistent underlying trigger, or the continued existence of events with that Reason.  Events should be treated as informative, best-effort, supplemental data.",
		MarkdownDescription: "Event is a report of an event somewhere in the cluster. It generally denotes some state change in the system. Events have a limited retention time and triggers and messages may evolve with time.  Event consumers should not rely on the timing of an event with a given Reason reflecting a consistent underlying trigger, or the continued existence of events with that Reason.  Events should be treated as informative, best-effort, supplemental data.",
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

			"action": {
				Description:         "action is what action was taken/failed regarding to the regarding object. It is machine-readable. This field cannot be empty for new Events and it can have at most 128 characters.",
				MarkdownDescription: "action is what action was taken/failed regarding to the regarding object. It is machine-readable. This field cannot be empty for new Events and it can have at most 128 characters.",

				Type: types.StringType,

				Required: false,
				Optional: true,
				Computed: false,
			},

			"deprecated_count": {
				Description:         "deprecatedCount is the deprecated field assuring backward compatibility with core.v1 Event type.",
				MarkdownDescription: "deprecatedCount is the deprecated field assuring backward compatibility with core.v1 Event type.",

				Type: types.Int64Type,

				Required: false,
				Optional: true,
				Computed: false,
			},

			"deprecated_first_timestamp": {
				Description:         "Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.",
				MarkdownDescription: "Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.",

				Type: types.StringType,

				Required: false,
				Optional: true,
				Computed: false,

				Validators: []tfsdk.AttributeValidator{

					validators.DateTime64Validator(),
				},
			},

			"deprecated_last_timestamp": {
				Description:         "Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.",
				MarkdownDescription: "Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.",

				Type: types.StringType,

				Required: false,
				Optional: true,
				Computed: false,

				Validators: []tfsdk.AttributeValidator{

					validators.DateTime64Validator(),
				},
			},

			"deprecated_source": {
				Description:         "EventSource contains information for an event.",
				MarkdownDescription: "EventSource contains information for an event.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"component": {
						Description:         "Component from which the event is generated.",
						MarkdownDescription: "Component from which the event is generated.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"host": {
						Description:         "Node name on which the event is generated.",
						MarkdownDescription: "Node name on which the event is generated.",

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

			"event_time": {
				Description:         "MicroTime is version of Time with microsecond level precision.",
				MarkdownDescription: "MicroTime is version of Time with microsecond level precision.",

				Type: types.StringType,

				Required: true,
				Optional: false,
				Computed: false,

				Validators: []tfsdk.AttributeValidator{

					validators.DateTime64Validator(),
				},
			},

			"note": {
				Description:         "note is a human-readable description of the status of this operation. Maximal length of the note is 1kB, but libraries should be prepared to handle values up to 64kB.",
				MarkdownDescription: "note is a human-readable description of the status of this operation. Maximal length of the note is 1kB, but libraries should be prepared to handle values up to 64kB.",

				Type: types.StringType,

				Required: false,
				Optional: true,
				Computed: false,
			},

			"reason": {
				Description:         "reason is why the action was taken. It is human-readable. This field cannot be empty for new Events and it can have at most 128 characters.",
				MarkdownDescription: "reason is why the action was taken. It is human-readable. This field cannot be empty for new Events and it can have at most 128 characters.",

				Type: types.StringType,

				Required: false,
				Optional: true,
				Computed: false,
			},

			"regarding": {
				Description:         "ObjectReference contains enough information to let you inspect or modify the referred object.",
				MarkdownDescription: "ObjectReference contains enough information to let you inspect or modify the referred object.",

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
						Description:         "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object.",
						MarkdownDescription: "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object.",

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

			"related": {
				Description:         "ObjectReference contains enough information to let you inspect or modify the referred object.",
				MarkdownDescription: "ObjectReference contains enough information to let you inspect or modify the referred object.",

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
						Description:         "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object.",
						MarkdownDescription: "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object.",

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

			"reporting_controller": {
				Description:         "reportingController is the name of the controller that emitted this Event, e.g. 'kubernetes.io/kubelet'. This field cannot be empty for new Events.",
				MarkdownDescription: "reportingController is the name of the controller that emitted this Event, e.g. 'kubernetes.io/kubelet'. This field cannot be empty for new Events.",

				Type: types.StringType,

				Required: false,
				Optional: true,
				Computed: false,
			},

			"reporting_instance": {
				Description:         "reportingInstance is the ID of the controller instance, e.g. 'kubelet-xyzf'. This field cannot be empty for new Events and it can have at most 128 characters.",
				MarkdownDescription: "reportingInstance is the ID of the controller instance, e.g. 'kubelet-xyzf'. This field cannot be empty for new Events and it can have at most 128 characters.",

				Type: types.StringType,

				Required: false,
				Optional: true,
				Computed: false,
			},

			"series": {
				Description:         "EventSeries contain information on series of events, i.e. thing that was/is happening continuously for some time. How often to update the EventSeries is up to the event reporters. The default event reporter in 'k8s.io/client-go/tools/events/event_broadcaster.go' shows how this struct is updated on heartbeats and can guide customized reporter implementations.",
				MarkdownDescription: "EventSeries contain information on series of events, i.e. thing that was/is happening continuously for some time. How often to update the EventSeries is up to the event reporters. The default event reporter in 'k8s.io/client-go/tools/events/event_broadcaster.go' shows how this struct is updated on heartbeats and can guide customized reporter implementations.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"count": {
						Description:         "count is the number of occurrences in this series up to the last heartbeat time.",
						MarkdownDescription: "count is the number of occurrences in this series up to the last heartbeat time.",

						Type: types.Int64Type,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"last_observed_time": {
						Description:         "MicroTime is version of Time with microsecond level precision.",
						MarkdownDescription: "MicroTime is version of Time with microsecond level precision.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							validators.DateTime64Validator(),
						},
					},
				}),

				Required: false,
				Optional: true,
				Computed: false,
			},

			"type": {
				Description:         "type is the type of this event (Normal, Warning), new types could be added in the future. It is machine-readable. This field cannot be empty for new Events.",
				MarkdownDescription: "type is the type of this event (Normal, Warning), new types could be added in the future. It is machine-readable. This field cannot be empty for new Events.",

				Type: types.StringType,

				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}, nil
}

func (r *EventsK8SIoEventV1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_events_k8s_io_event_v1")

	var state EventsK8SIoEventV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel EventsK8SIoEventV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("events.k8s.io/v1")
	goModel.Kind = utilities.Ptr("Event")

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

func (r *EventsK8SIoEventV1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_events_k8s_io_event_v1")
	// NO-OP: All data is already in Terraform state
}

func (r *EventsK8SIoEventV1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_events_k8s_io_event_v1")

	var state EventsK8SIoEventV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel EventsK8SIoEventV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("events.k8s.io/v1")
	goModel.Kind = utilities.Ptr("Event")

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

func (r *EventsK8SIoEventV1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_events_k8s_io_event_v1")
	// NO-OP: Terraform removes the state automatically for us
}
