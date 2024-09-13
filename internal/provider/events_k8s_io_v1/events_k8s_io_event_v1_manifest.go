/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package events_k8s_io_v1

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
	_ datasource.DataSource = &EventsK8SIoEventV1Manifest{}
)

func NewEventsK8SIoEventV1Manifest() datasource.DataSource {
	return &EventsK8SIoEventV1Manifest{}
}

type EventsK8SIoEventV1Manifest struct{}

type EventsK8SIoEventV1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Namespace   string            `tfsdk:"namespace" json:"namespace"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Action                   *string `tfsdk:"action" json:"action,omitempty"`
	DeprecatedCount          *int64  `tfsdk:"deprecated_count" json:"deprecatedCount,omitempty"`
	DeprecatedFirstTimestamp *string `tfsdk:"deprecated_first_timestamp" json:"deprecatedFirstTimestamp,omitempty"`
	DeprecatedLastTimestamp  *string `tfsdk:"deprecated_last_timestamp" json:"deprecatedLastTimestamp,omitempty"`
	DeprecatedSource         *struct {
		Component *string `tfsdk:"component" json:"component,omitempty"`
		Host      *string `tfsdk:"host" json:"host,omitempty"`
	} `tfsdk:"deprecated_source" json:"deprecatedSource,omitempty"`
	EventTime *string `tfsdk:"event_time" json:"eventTime,omitempty"`
	Note      *string `tfsdk:"note" json:"note,omitempty"`
	Reason    *string `tfsdk:"reason" json:"reason,omitempty"`
	Regarding *struct {
		ApiVersion      *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
		FieldPath       *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
		Kind            *string `tfsdk:"kind" json:"kind,omitempty"`
		Name            *string `tfsdk:"name" json:"name,omitempty"`
		Namespace       *string `tfsdk:"namespace" json:"namespace,omitempty"`
		ResourceVersion *string `tfsdk:"resource_version" json:"resourceVersion,omitempty"`
		Uid             *string `tfsdk:"uid" json:"uid,omitempty"`
	} `tfsdk:"regarding" json:"regarding,omitempty"`
	Related *struct {
		ApiVersion      *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
		FieldPath       *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
		Kind            *string `tfsdk:"kind" json:"kind,omitempty"`
		Name            *string `tfsdk:"name" json:"name,omitempty"`
		Namespace       *string `tfsdk:"namespace" json:"namespace,omitempty"`
		ResourceVersion *string `tfsdk:"resource_version" json:"resourceVersion,omitempty"`
		Uid             *string `tfsdk:"uid" json:"uid,omitempty"`
	} `tfsdk:"related" json:"related,omitempty"`
	ReportingController *string `tfsdk:"reporting_controller" json:"reportingController,omitempty"`
	ReportingInstance   *string `tfsdk:"reporting_instance" json:"reportingInstance,omitempty"`
	Series              *struct {
		Count            *int64  `tfsdk:"count" json:"count,omitempty"`
		LastObservedTime *string `tfsdk:"last_observed_time" json:"lastObservedTime,omitempty"`
	} `tfsdk:"series" json:"series,omitempty"`
	Type *string `tfsdk:"type" json:"type,omitempty"`
}

func (r *EventsK8SIoEventV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_events_k8s_io_event_v1_manifest"
}

func (r *EventsK8SIoEventV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Event is a report of an event somewhere in the cluster. It generally denotes some state change in the system. Events have a limited retention time and triggers and messages may evolve with time. Event consumers should not rely on the timing of an event with a given Reason reflecting a consistent underlying trigger, or the continued existence of events with that Reason. Events should be treated as informative, best-effort, supplemental data.",
		MarkdownDescription: "Event is a report of an event somewhere in the cluster. It generally denotes some state change in the system. Events have a limited retention time and triggers and messages may evolve with time. Event consumers should not rely on the timing of an event with a given Reason reflecting a consistent underlying trigger, or the continued existence of events with that Reason. Events should be treated as informative, best-effort, supplemental data.",
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

			"action": schema.StringAttribute{
				Description:         "action is what action was taken/failed regarding to the regarding object. It is machine-readable. This field cannot be empty for new Events and it can have at most 128 characters.",
				MarkdownDescription: "action is what action was taken/failed regarding to the regarding object. It is machine-readable. This field cannot be empty for new Events and it can have at most 128 characters.",
				Required:            false,
				Optional:            true,
				Computed:            false,
			},

			"deprecated_count": schema.Int64Attribute{
				Description:         "deprecatedCount is the deprecated field assuring backward compatibility with core.v1 Event type.",
				MarkdownDescription: "deprecatedCount is the deprecated field assuring backward compatibility with core.v1 Event type.",
				Required:            false,
				Optional:            true,
				Computed:            false,
			},

			"deprecated_first_timestamp": schema.StringAttribute{
				Description:         "Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON. Wrappers are provided for many of the factory methods that the time package offers.",
				MarkdownDescription: "Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON. Wrappers are provided for many of the factory methods that the time package offers.",
				Required:            false,
				Optional:            true,
				Computed:            false,
				Validators: []validator.String{
					validators.DateTime64Validator(),
				},
			},

			"deprecated_last_timestamp": schema.StringAttribute{
				Description:         "Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON. Wrappers are provided for many of the factory methods that the time package offers.",
				MarkdownDescription: "Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON. Wrappers are provided for many of the factory methods that the time package offers.",
				Required:            false,
				Optional:            true,
				Computed:            false,
				Validators: []validator.String{
					validators.DateTime64Validator(),
				},
			},

			"deprecated_source": schema.SingleNestedAttribute{
				Description:         "EventSource contains information for an event.",
				MarkdownDescription: "EventSource contains information for an event.",
				Attributes: map[string]schema.Attribute{
					"component": schema.StringAttribute{
						Description:         "Component from which the event is generated.",
						MarkdownDescription: "Component from which the event is generated.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"host": schema.StringAttribute{
						Description:         "Node name on which the event is generated.",
						MarkdownDescription: "Node name on which the event is generated.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},
				},
				Required: false,
				Optional: true,
				Computed: false,
			},

			"event_time": schema.StringAttribute{
				Description:         "MicroTime is version of Time with microsecond level precision.",
				MarkdownDescription: "MicroTime is version of Time with microsecond level precision.",
				Required:            true,
				Optional:            false,
				Computed:            false,
				Validators: []validator.String{
					validators.DateTime64Validator(),
				},
			},

			"note": schema.StringAttribute{
				Description:         "note is a human-readable description of the status of this operation. Maximal length of the note is 1kB, but libraries should be prepared to handle values up to 64kB.",
				MarkdownDescription: "note is a human-readable description of the status of this operation. Maximal length of the note is 1kB, but libraries should be prepared to handle values up to 64kB.",
				Required:            false,
				Optional:            true,
				Computed:            false,
			},

			"reason": schema.StringAttribute{
				Description:         "reason is why the action was taken. It is human-readable. This field cannot be empty for new Events and it can have at most 128 characters.",
				MarkdownDescription: "reason is why the action was taken. It is human-readable. This field cannot be empty for new Events and it can have at most 128 characters.",
				Required:            false,
				Optional:            true,
				Computed:            false,
			},

			"regarding": schema.SingleNestedAttribute{
				Description:         "ObjectReference contains enough information to let you inspect or modify the referred object.",
				MarkdownDescription: "ObjectReference contains enough information to let you inspect or modify the referred object.",
				Attributes: map[string]schema.Attribute{
					"api_version": schema.StringAttribute{
						Description:         "API version of the referent.",
						MarkdownDescription: "API version of the referent.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"field_path": schema.StringAttribute{
						Description:         "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object.",
						MarkdownDescription: "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"kind": schema.StringAttribute{
						Description:         "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
						MarkdownDescription: "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"name": schema.StringAttribute{
						Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
						MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"namespace": schema.StringAttribute{
						Description:         "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
						MarkdownDescription: "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"resource_version": schema.StringAttribute{
						Description:         "Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
						MarkdownDescription: "Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"uid": schema.StringAttribute{
						Description:         "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
						MarkdownDescription: "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},
				},
				Required: false,
				Optional: true,
				Computed: false,
			},

			"related": schema.SingleNestedAttribute{
				Description:         "ObjectReference contains enough information to let you inspect or modify the referred object.",
				MarkdownDescription: "ObjectReference contains enough information to let you inspect or modify the referred object.",
				Attributes: map[string]schema.Attribute{
					"api_version": schema.StringAttribute{
						Description:         "API version of the referent.",
						MarkdownDescription: "API version of the referent.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"field_path": schema.StringAttribute{
						Description:         "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object.",
						MarkdownDescription: "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"kind": schema.StringAttribute{
						Description:         "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
						MarkdownDescription: "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"name": schema.StringAttribute{
						Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
						MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"namespace": schema.StringAttribute{
						Description:         "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
						MarkdownDescription: "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"resource_version": schema.StringAttribute{
						Description:         "Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
						MarkdownDescription: "Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"uid": schema.StringAttribute{
						Description:         "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
						MarkdownDescription: "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},
				},
				Required: false,
				Optional: true,
				Computed: false,
			},

			"reporting_controller": schema.StringAttribute{
				Description:         "reportingController is the name of the controller that emitted this Event, e.g. 'kubernetes.io/kubelet'. This field cannot be empty for new Events.",
				MarkdownDescription: "reportingController is the name of the controller that emitted this Event, e.g. 'kubernetes.io/kubelet'. This field cannot be empty for new Events.",
				Required:            false,
				Optional:            true,
				Computed:            false,
			},

			"reporting_instance": schema.StringAttribute{
				Description:         "reportingInstance is the ID of the controller instance, e.g. 'kubelet-xyzf'. This field cannot be empty for new Events and it can have at most 128 characters.",
				MarkdownDescription: "reportingInstance is the ID of the controller instance, e.g. 'kubelet-xyzf'. This field cannot be empty for new Events and it can have at most 128 characters.",
				Required:            false,
				Optional:            true,
				Computed:            false,
			},

			"series": schema.SingleNestedAttribute{
				Description:         "EventSeries contain information on series of events, i.e. thing that was/is happening continuously for some time. How often to update the EventSeries is up to the event reporters. The default event reporter in 'k8s.io/client-go/tools/events/event_broadcaster.go' shows how this struct is updated on heartbeats and can guide customized reporter implementations.",
				MarkdownDescription: "EventSeries contain information on series of events, i.e. thing that was/is happening continuously for some time. How often to update the EventSeries is up to the event reporters. The default event reporter in 'k8s.io/client-go/tools/events/event_broadcaster.go' shows how this struct is updated on heartbeats and can guide customized reporter implementations.",
				Attributes: map[string]schema.Attribute{
					"count": schema.Int64Attribute{
						Description:         "count is the number of occurrences in this series up to the last heartbeat time.",
						MarkdownDescription: "count is the number of occurrences in this series up to the last heartbeat time.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"last_observed_time": schema.StringAttribute{
						Description:         "MicroTime is version of Time with microsecond level precision.",
						MarkdownDescription: "MicroTime is version of Time with microsecond level precision.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							validators.DateTime64Validator(),
						},
					},
				},
				Required: false,
				Optional: true,
				Computed: false,
			},

			"type": schema.StringAttribute{
				Description:         "type is the type of this event (Normal, Warning), new types could be added in the future. It is machine-readable. This field cannot be empty for new Events.",
				MarkdownDescription: "type is the type of this event (Normal, Warning), new types could be added in the future. It is machine-readable. This field cannot be empty for new Events.",
				Required:            false,
				Optional:            true,
				Computed:            false,
			},
		},
	}
}

func (r *EventsK8SIoEventV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_events_k8s_io_event_v1_manifest")

	var model EventsK8SIoEventV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("events.k8s.io/v1")
	model.Kind = pointer.String("Event")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
