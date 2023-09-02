/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package kafka_strimzi_io_v1beta2

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
	_ datasource.DataSource              = &KafkaStrimziIoKafkaRebalanceV1Beta2DataSource{}
	_ datasource.DataSourceWithConfigure = &KafkaStrimziIoKafkaRebalanceV1Beta2DataSource{}
)

func NewKafkaStrimziIoKafkaRebalanceV1Beta2DataSource() datasource.DataSource {
	return &KafkaStrimziIoKafkaRebalanceV1Beta2DataSource{}
}

type KafkaStrimziIoKafkaRebalanceV1Beta2DataSource struct {
	kubernetesClient dynamic.Interface
}

type KafkaStrimziIoKafkaRebalanceV1Beta2DataSourceData struct {
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

func (r *KafkaStrimziIoKafkaRebalanceV1Beta2DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_kafka_strimzi_io_kafka_rebalance_v1beta2"
}

func (r *KafkaStrimziIoKafkaRebalanceV1Beta2DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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
				Description:         "The specification of the Kafka rebalance.",
				MarkdownDescription: "The specification of the Kafka rebalance.",
				Attributes: map[string]schema.Attribute{
					"brokers": schema.ListAttribute{
						Description:         "The list of newly added brokers in case of scaling up or the ones to be removed in case of scaling down to use for rebalancing. This list can be used only with rebalancing mode 'add-brokers' and 'removed-brokers'. It is ignored with 'full' mode.",
						MarkdownDescription: "The list of newly added brokers in case of scaling up or the ones to be removed in case of scaling down to use for rebalancing. This list can be used only with rebalancing mode 'add-brokers' and 'removed-brokers'. It is ignored with 'full' mode.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"concurrent_intra_broker_partition_movements": schema.Int64Attribute{
						Description:         "The upper bound of ongoing partition replica movements between disks within each broker. Default is 2.",
						MarkdownDescription: "The upper bound of ongoing partition replica movements between disks within each broker. Default is 2.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"concurrent_leader_movements": schema.Int64Attribute{
						Description:         "The upper bound of ongoing partition leadership movements. Default is 1000.",
						MarkdownDescription: "The upper bound of ongoing partition leadership movements. Default is 1000.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"concurrent_partition_movements_per_broker": schema.Int64Attribute{
						Description:         "The upper bound of ongoing partition replica movements going into/out of each broker. Default is 5.",
						MarkdownDescription: "The upper bound of ongoing partition replica movements going into/out of each broker. Default is 5.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"excluded_topics": schema.StringAttribute{
						Description:         "A regular expression where any matching topics will be excluded from the calculation of optimization proposals. This expression will be parsed by the java.util.regex.Pattern class; for more information on the supported format consult the documentation for that class.",
						MarkdownDescription: "A regular expression where any matching topics will be excluded from the calculation of optimization proposals. This expression will be parsed by the java.util.regex.Pattern class; for more information on the supported format consult the documentation for that class.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"goals": schema.ListAttribute{
						Description:         "A list of goals, ordered by decreasing priority, to use for generating and executing the rebalance proposal. The supported goals are available at https://github.com/linkedin/cruise-control#goals. If an empty goals list is provided, the goals declared in the default.goals Cruise Control configuration parameter are used.",
						MarkdownDescription: "A list of goals, ordered by decreasing priority, to use for generating and executing the rebalance proposal. The supported goals are available at https://github.com/linkedin/cruise-control#goals. If an empty goals list is provided, the goals declared in the default.goals Cruise Control configuration parameter are used.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"mode": schema.StringAttribute{
						Description:         "Mode to run the rebalancing. The supported modes are 'full', 'add-brokers', 'remove-brokers'.If not specified, the 'full' mode is used by default. * 'full' mode runs the rebalancing across all the brokers in the cluster.* 'add-brokers' mode can be used after scaling up the cluster to move some replicas to the newly added brokers.* 'remove-brokers' mode can be used before scaling down the cluster to move replicas out of the brokers to be removed.",
						MarkdownDescription: "Mode to run the rebalancing. The supported modes are 'full', 'add-brokers', 'remove-brokers'.If not specified, the 'full' mode is used by default. * 'full' mode runs the rebalancing across all the brokers in the cluster.* 'add-brokers' mode can be used after scaling up the cluster to move some replicas to the newly added brokers.* 'remove-brokers' mode can be used before scaling down the cluster to move replicas out of the brokers to be removed.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"rebalance_disk": schema.BoolAttribute{
						Description:         "Enables intra-broker disk balancing, which balances disk space utilization between disks on the same broker. Only applies to Kafka deployments that use JBOD storage with multiple disks. When enabled, inter-broker balancing is disabled. Default is false.",
						MarkdownDescription: "Enables intra-broker disk balancing, which balances disk space utilization between disks on the same broker. Only applies to Kafka deployments that use JBOD storage with multiple disks. When enabled, inter-broker balancing is disabled. Default is false.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"replica_movement_strategies": schema.ListAttribute{
						Description:         "A list of strategy class names used to determine the execution order for the replica movements in the generated optimization proposal. By default BaseReplicaMovementStrategy is used, which will execute the replica movements in the order that they were generated.",
						MarkdownDescription: "A list of strategy class names used to determine the execution order for the replica movements in the generated optimization proposal. By default BaseReplicaMovementStrategy is used, which will execute the replica movements in the order that they were generated.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"replication_throttle": schema.Int64Attribute{
						Description:         "The upper bound, in bytes per second, on the bandwidth used to move replicas. There is no limit by default.",
						MarkdownDescription: "The upper bound, in bytes per second, on the bandwidth used to move replicas. There is no limit by default.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"skip_hard_goal_check": schema.BoolAttribute{
						Description:         "Whether to allow the hard goals specified in the Kafka CR to be skipped in optimization proposal generation. This can be useful when some of those hard goals are preventing a balance solution being found. Default is false.",
						MarkdownDescription: "Whether to allow the hard goals specified in the Kafka CR to be skipped in optimization proposal generation. This can be useful when some of those hard goals are preventing a balance solution being found. Default is false.",
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

func (r *KafkaStrimziIoKafkaRebalanceV1Beta2DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *KafkaStrimziIoKafkaRebalanceV1Beta2DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_kafka_strimzi_io_kafka_rebalance_v1beta2")

	var data KafkaStrimziIoKafkaRebalanceV1Beta2DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "kafka.strimzi.io", Version: "v1beta2", Resource: "KafkaRebalance"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to GET resource",
			"An unexpected error occurred while reading the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"GET Error: "+err.Error(),
		)
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

	var readResponse KafkaStrimziIoKafkaRebalanceV1Beta2DataSourceData
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
	data.ApiVersion = pointer.String("kafka.strimzi.io/v1beta2")
	data.Kind = pointer.String("KafkaRebalance")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
