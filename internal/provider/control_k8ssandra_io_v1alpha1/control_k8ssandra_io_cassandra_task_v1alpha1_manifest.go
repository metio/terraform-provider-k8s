/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package control_k8ssandra_io_v1alpha1

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
	_ datasource.DataSource = &ControlK8SsandraIoCassandraTaskV1Alpha1Manifest{}
)

func NewControlK8SsandraIoCassandraTaskV1Alpha1Manifest() datasource.DataSource {
	return &ControlK8SsandraIoCassandraTaskV1Alpha1Manifest{}
}

type ControlK8SsandraIoCassandraTaskV1Alpha1Manifest struct{}

type ControlK8SsandraIoCassandraTaskV1Alpha1ManifestData struct {
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
		ConcurrencyPolicy *string `tfsdk:"concurrency_policy" json:"concurrencyPolicy,omitempty"`
		Datacenter        *struct {
			ApiVersion      *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
			FieldPath       *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
			Kind            *string `tfsdk:"kind" json:"kind,omitempty"`
			Name            *string `tfsdk:"name" json:"name,omitempty"`
			Namespace       *string `tfsdk:"namespace" json:"namespace,omitempty"`
			ResourceVersion *string `tfsdk:"resource_version" json:"resourceVersion,omitempty"`
			Uid             *string `tfsdk:"uid" json:"uid,omitempty"`
		} `tfsdk:"datacenter" json:"datacenter,omitempty"`
		Jobs *[]struct {
			Args *struct {
				End_token         *string            `tfsdk:"end_token" json:"end_token,omitempty"`
				Jobs              *int64             `tfsdk:"jobs" json:"jobs,omitempty"`
				Keyspace_name     *string            `tfsdk:"keyspace_name" json:"keyspace_name,omitempty"`
				New_tokens        *map[string]string `tfsdk:"new_tokens" json:"new_tokens,omitempty"`
				No_snapshot       *bool              `tfsdk:"no_snapshot" json:"no_snapshot,omitempty"`
				No_validate       *bool              `tfsdk:"no_validate" json:"no_validate,omitempty"`
				Pod_name          *string            `tfsdk:"pod_name" json:"pod_name,omitempty"`
				Rack              *string            `tfsdk:"rack" json:"rack,omitempty"`
				Skip_corrupted    *bool              `tfsdk:"skip_corrupted" json:"skip_corrupted,omitempty"`
				Source_datacenter *string            `tfsdk:"source_datacenter" json:"source_datacenter,omitempty"`
				Split_output      *bool              `tfsdk:"split_output" json:"split_output,omitempty"`
				Start_token       *string            `tfsdk:"start_token" json:"start_token,omitempty"`
				Tables            *[]string          `tfsdk:"tables" json:"tables,omitempty"`
			} `tfsdk:"args" json:"args,omitempty"`
			Command *string `tfsdk:"command" json:"command,omitempty"`
			Name    *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"jobs" json:"jobs,omitempty"`
		RestartPolicy           *string `tfsdk:"restart_policy" json:"restartPolicy,omitempty"`
		ScheduledTime           *string `tfsdk:"scheduled_time" json:"scheduledTime,omitempty"`
		TtlSecondsAfterFinished *int64  `tfsdk:"ttl_seconds_after_finished" json:"ttlSecondsAfterFinished,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ControlK8SsandraIoCassandraTaskV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_control_k8ssandra_io_cassandra_task_v1alpha1_manifest"
}

func (r *ControlK8SsandraIoCassandraTaskV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "CassandraTask is the Schema for the cassandrajobs API",
		MarkdownDescription: "CassandraTask is the Schema for the cassandrajobs API",
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
				Description:         "CassandraTaskSpec defines the desired state of CassandraTask",
				MarkdownDescription: "CassandraTaskSpec defines the desired state of CassandraTask",
				Attributes: map[string]schema.Attribute{
					"concurrency_policy": schema.StringAttribute{
						Description:         "Specifics if this task can be run concurrently with other active tasks. Valid values are:- 'Allow': allows multiple Tasks to run concurrently on Cassandra cluster- 'Forbid' (default): only a single task is executed at onceThe 'Allow' property is only valid if all the other active Tasks have 'Allow' as well.",
						MarkdownDescription: "Specifics if this task can be run concurrently with other active tasks. Valid values are:- 'Allow': allows multiple Tasks to run concurrently on Cassandra cluster- 'Forbid' (default): only a single task is executed at onceThe 'Allow' property is only valid if all the other active Tasks have 'Allow' as well.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"datacenter": schema.SingleNestedAttribute{
						Description:         "Which datacenter this task is targetting. Note, this must be a datacenter which the current cass-operatorcan access",
						MarkdownDescription: "Which datacenter this task is targetting. Note, this must be a datacenter which the current cass-operatorcan access",
						Attributes: map[string]schema.Attribute{
							"api_version": schema.StringAttribute{
								Description:         "API version of the referent.",
								MarkdownDescription: "API version of the referent.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"field_path": schema.StringAttribute{
								Description:         "If referring to a piece of an object instead of an entire object, this stringshould contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2].For example, if the object reference is to a container within a pod, this would take on a value like:'spec.containers{name}' (where 'name' refers to the name of the container that triggeredthe event) or if no container name is specified 'spec.containers[2]' (container withindex 2 in this pod). This syntax is chosen only to have some well-defined way ofreferencing a part of an object.TODO: this design is not final and this field is subject to change in the future.",
								MarkdownDescription: "If referring to a piece of an object instead of an entire object, this stringshould contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2].For example, if the object reference is to a container within a pod, this would take on a value like:'spec.containers{name}' (where 'name' refers to the name of the container that triggeredthe event) or if no container name is specified 'spec.containers[2]' (container withindex 2 in this pod). This syntax is chosen only to have some well-defined way ofreferencing a part of an object.TODO: this design is not final and this field is subject to change in the future.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kind": schema.StringAttribute{
								Description:         "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
								MarkdownDescription: "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
								MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"namespace": schema.StringAttribute{
								Description:         "Namespace of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
								MarkdownDescription: "Namespace of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"resource_version": schema.StringAttribute{
								Description:         "Specific resourceVersion to which this reference is made, if any.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
								MarkdownDescription: "Specific resourceVersion to which this reference is made, if any.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"uid": schema.StringAttribute{
								Description:         "UID of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
								MarkdownDescription: "UID of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"jobs": schema.ListNestedAttribute{
						Description:         "Jobs defines the jobs this task will execute (and their order)",
						MarkdownDescription: "Jobs defines the jobs this task will execute (and their order)",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"args": schema.SingleNestedAttribute{
									Description:         "Arguments are additional parameters for the command",
									MarkdownDescription: "Arguments are additional parameters for the command",
									Attributes: map[string]schema.Attribute{
										"end_token": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"jobs": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"keyspace_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"new_tokens": schema.MapAttribute{
											Description:         "NewTokens is a map of pod names to their newly-assigned tokens. Required for the movecommand, ignored otherwise. Pods referenced in this map must exist; any existing pod notreferenced in this map will not be moved.",
											MarkdownDescription: "NewTokens is a map of pod names to their newly-assigned tokens. Required for the movecommand, ignored otherwise. Pods referenced in this map must exist; any existing pod notreferenced in this map will not be moved.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"no_snapshot": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"no_validate": schema.BoolAttribute{
											Description:         "Scrub arguments",
											MarkdownDescription: "Scrub arguments",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"pod_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"rack": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"skip_corrupted": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"source_datacenter": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"split_output": schema.BoolAttribute{
											Description:         "Compaction arguments",
											MarkdownDescription: "Compaction arguments",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"start_token": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"tables": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"command": schema.StringAttribute{
									Description:         "Command defines what is run against Cassandra pods",
									MarkdownDescription: "Command defines what is run against Cassandra pods",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"restart_policy": schema.StringAttribute{
						Description:         "RestartPolicy indicates the behavior n case of failure. Default is Never.",
						MarkdownDescription: "RestartPolicy indicates the behavior n case of failure. Default is Never.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"scheduled_time": schema.StringAttribute{
						Description:         "ScheduledTime indicates the earliest possible time this task is executed. This does not necessarilyequal to the time it is actually executed (if other tasks are blocking for example). If not set,the task will be executed immediately.",
						MarkdownDescription: "ScheduledTime indicates the earliest possible time this task is executed. This does not necessarilyequal to the time it is actually executed (if other tasks are blocking for example). If not set,the task will be executed immediately.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							validators.DateTime64Validator(),
						},
					},

					"ttl_seconds_after_finished": schema.Int64Attribute{
						Description:         "TTLSecondsAfterFinished defines how long the completed job will kept before being cleaned up. If set to 0the task will not be cleaned up by the cass-operator. If unset, the default time (86400s) is used.",
						MarkdownDescription: "TTLSecondsAfterFinished defines how long the completed job will kept before being cleaned up. If set to 0the task will not be cleaned up by the cass-operator. If unset, the default time (86400s) is used.",
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

func (r *ControlK8SsandraIoCassandraTaskV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_control_k8ssandra_io_cassandra_task_v1alpha1_manifest")

	var model ControlK8SsandraIoCassandraTaskV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("control.k8ssandra.io/v1alpha1")
	model.Kind = pointer.String("CassandraTask")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
