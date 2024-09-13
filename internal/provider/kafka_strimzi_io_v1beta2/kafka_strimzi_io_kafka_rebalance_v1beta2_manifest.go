/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package kafka_strimzi_io_v1beta2

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
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
	_ datasource.DataSource = &KafkaStrimziIoKafkaRebalanceV1Beta2Manifest{}
)

func NewKafkaStrimziIoKafkaRebalanceV1Beta2Manifest() datasource.DataSource {
	return &KafkaStrimziIoKafkaRebalanceV1Beta2Manifest{}
}

type KafkaStrimziIoKafkaRebalanceV1Beta2Manifest struct{}

type KafkaStrimziIoKafkaRebalanceV1Beta2ManifestData struct {
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
		Brokers                                 *[]string `tfsdk:"brokers" json:"brokers,omitempty"`
		ConcurrentIntraBrokerPartitionMovements *int64    `tfsdk:"concurrent_intra_broker_partition_movements" json:"concurrentIntraBrokerPartitionMovements,omitempty"`
		ConcurrentLeaderMovements               *int64    `tfsdk:"concurrent_leader_movements" json:"concurrentLeaderMovements,omitempty"`
		ConcurrentPartitionMovementsPerBroker   *int64    `tfsdk:"concurrent_partition_movements_per_broker" json:"concurrentPartitionMovementsPerBroker,omitempty"`
		ExcludedTopics                          *string   `tfsdk:"excluded_topics" json:"excludedTopics,omitempty"`
		Goals                                   *[]string `tfsdk:"goals" json:"goals,omitempty"`
		Mode                                    *string   `tfsdk:"mode" json:"mode,omitempty"`
		RebalanceDisk                           *bool     `tfsdk:"rebalance_disk" json:"rebalanceDisk,omitempty"`
		ReplicaMovementStrategies               *[]string `tfsdk:"replica_movement_strategies" json:"replicaMovementStrategies,omitempty"`
		ReplicationThrottle                     *int64    `tfsdk:"replication_throttle" json:"replicationThrottle,omitempty"`
		SkipHardGoalCheck                       *bool     `tfsdk:"skip_hard_goal_check" json:"skipHardGoalCheck,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *KafkaStrimziIoKafkaRebalanceV1Beta2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_kafka_strimzi_io_kafka_rebalance_v1beta2_manifest"
}

func (r *KafkaStrimziIoKafkaRebalanceV1Beta2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "",
		MarkdownDescription: "",
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
				Description:         "The specification of the Kafka rebalance.",
				MarkdownDescription: "The specification of the Kafka rebalance.",
				Attributes: map[string]schema.Attribute{
					"brokers": schema.ListAttribute{
						Description:         "The list of newly added brokers in case of scaling up or the ones to be removed in case of scaling down to use for rebalancing. This list can be used only with rebalancing mode 'add-brokers' and 'removed-brokers'. It is ignored with 'full' mode.",
						MarkdownDescription: "The list of newly added brokers in case of scaling up or the ones to be removed in case of scaling down to use for rebalancing. This list can be used only with rebalancing mode 'add-brokers' and 'removed-brokers'. It is ignored with 'full' mode.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"concurrent_intra_broker_partition_movements": schema.Int64Attribute{
						Description:         "The upper bound of ongoing partition replica movements between disks within each broker. Default is 2.",
						MarkdownDescription: "The upper bound of ongoing partition replica movements between disks within each broker. Default is 2.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
						},
					},

					"concurrent_leader_movements": schema.Int64Attribute{
						Description:         "The upper bound of ongoing partition leadership movements. Default is 1000.",
						MarkdownDescription: "The upper bound of ongoing partition leadership movements. Default is 1000.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
						},
					},

					"concurrent_partition_movements_per_broker": schema.Int64Attribute{
						Description:         "The upper bound of ongoing partition replica movements going into/out of each broker. Default is 5.",
						MarkdownDescription: "The upper bound of ongoing partition replica movements going into/out of each broker. Default is 5.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
						},
					},

					"excluded_topics": schema.StringAttribute{
						Description:         "A regular expression where any matching topics will be excluded from the calculation of optimization proposals. This expression will be parsed by the java.util.regex.Pattern class; for more information on the supported format consult the documentation for that class.",
						MarkdownDescription: "A regular expression where any matching topics will be excluded from the calculation of optimization proposals. This expression will be parsed by the java.util.regex.Pattern class; for more information on the supported format consult the documentation for that class.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"goals": schema.ListAttribute{
						Description:         "A list of goals, ordered by decreasing priority, to use for generating and executing the rebalance proposal. The supported goals are available at https://github.com/linkedin/cruise-control#goals. If an empty goals list is provided, the goals declared in the default.goals Cruise Control configuration parameter are used.",
						MarkdownDescription: "A list of goals, ordered by decreasing priority, to use for generating and executing the rebalance proposal. The supported goals are available at https://github.com/linkedin/cruise-control#goals. If an empty goals list is provided, the goals declared in the default.goals Cruise Control configuration parameter are used.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"mode": schema.StringAttribute{
						Description:         "Mode to run the rebalancing. The supported modes are 'full', 'add-brokers', 'remove-brokers'. If not specified, the 'full' mode is used by default. * 'full' mode runs the rebalancing across all the brokers in the cluster. * 'add-brokers' mode can be used after scaling up the cluster to move some replicas to the newly added brokers. * 'remove-brokers' mode can be used before scaling down the cluster to move replicas out of the brokers to be removed. ",
						MarkdownDescription: "Mode to run the rebalancing. The supported modes are 'full', 'add-brokers', 'remove-brokers'. If not specified, the 'full' mode is used by default. * 'full' mode runs the rebalancing across all the brokers in the cluster. * 'add-brokers' mode can be used after scaling up the cluster to move some replicas to the newly added brokers. * 'remove-brokers' mode can be used before scaling down the cluster to move replicas out of the brokers to be removed. ",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("full", "add-brokers", "remove-brokers"),
						},
					},

					"rebalance_disk": schema.BoolAttribute{
						Description:         "Enables intra-broker disk balancing, which balances disk space utilization between disks on the same broker. Only applies to Kafka deployments that use JBOD storage with multiple disks. When enabled, inter-broker balancing is disabled. Default is false.",
						MarkdownDescription: "Enables intra-broker disk balancing, which balances disk space utilization between disks on the same broker. Only applies to Kafka deployments that use JBOD storage with multiple disks. When enabled, inter-broker balancing is disabled. Default is false.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"replica_movement_strategies": schema.ListAttribute{
						Description:         "A list of strategy class names used to determine the execution order for the replica movements in the generated optimization proposal. By default BaseReplicaMovementStrategy is used, which will execute the replica movements in the order that they were generated.",
						MarkdownDescription: "A list of strategy class names used to determine the execution order for the replica movements in the generated optimization proposal. By default BaseReplicaMovementStrategy is used, which will execute the replica movements in the order that they were generated.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"replication_throttle": schema.Int64Attribute{
						Description:         "The upper bound, in bytes per second, on the bandwidth used to move replicas. There is no limit by default.",
						MarkdownDescription: "The upper bound, in bytes per second, on the bandwidth used to move replicas. There is no limit by default.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
						},
					},

					"skip_hard_goal_check": schema.BoolAttribute{
						Description:         "Whether to allow the hard goals specified in the Kafka CR to be skipped in optimization proposal generation. This can be useful when some of those hard goals are preventing a balance solution being found. Default is false.",
						MarkdownDescription: "Whether to allow the hard goals specified in the Kafka CR to be skipped in optimization proposal generation. This can be useful when some of those hard goals are preventing a balance solution being found. Default is false.",
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

func (r *KafkaStrimziIoKafkaRebalanceV1Beta2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_kafka_strimzi_io_kafka_rebalance_v1beta2_manifest")

	var model KafkaStrimziIoKafkaRebalanceV1Beta2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("kafka.strimzi.io/v1beta2")
	model.Kind = pointer.String("KafkaRebalance")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
