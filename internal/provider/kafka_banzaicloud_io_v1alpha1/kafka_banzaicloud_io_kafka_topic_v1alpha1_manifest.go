/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package kafka_banzaicloud_io_v1alpha1

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
	_ datasource.DataSource = &KafkaBanzaicloudIoKafkaTopicV1Alpha1Manifest{}
)

func NewKafkaBanzaicloudIoKafkaTopicV1Alpha1Manifest() datasource.DataSource {
	return &KafkaBanzaicloudIoKafkaTopicV1Alpha1Manifest{}
}

type KafkaBanzaicloudIoKafkaTopicV1Alpha1Manifest struct{}

type KafkaBanzaicloudIoKafkaTopicV1Alpha1ManifestData struct {
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
		ClusterRef *struct {
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"cluster_ref" json:"clusterRef,omitempty"`
		Config            *map[string]string `tfsdk:"config" json:"config,omitempty"`
		Name              *string            `tfsdk:"name" json:"name,omitempty"`
		Partitions        *int64             `tfsdk:"partitions" json:"partitions,omitempty"`
		ReplicationFactor *int64             `tfsdk:"replication_factor" json:"replicationFactor,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *KafkaBanzaicloudIoKafkaTopicV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_kafka_banzaicloud_io_kafka_topic_v1alpha1_manifest"
}

func (r *KafkaBanzaicloudIoKafkaTopicV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "KafkaTopic is the Schema for the kafkatopics API",
		MarkdownDescription: "KafkaTopic is the Schema for the kafkatopics API",
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
				Description:         "KafkaTopicSpec defines the desired state of KafkaTopic",
				MarkdownDescription: "KafkaTopicSpec defines the desired state of KafkaTopic",
				Attributes: map[string]schema.Attribute{
					"cluster_ref": schema.SingleNestedAttribute{
						Description:         "ClusterReference states a reference to a cluster for topic/user provisioning",
						MarkdownDescription: "ClusterReference states a reference to a cluster for topic/user provisioning",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"namespace": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"config": schema.MapAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"name": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"partitions": schema.Int64Attribute{
						Description:         "Partitions defines the desired number of partitions; must be positive, or -1 to signify using the broker's default",
						MarkdownDescription: "Partitions defines the desired number of partitions; must be positive, or -1 to signify using the broker's default",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(-1),
						},
					},

					"replication_factor": schema.Int64Attribute{
						Description:         "ReplicationFactor defines the desired replication factor; must be positive, or -1 to signify using the broker's default",
						MarkdownDescription: "ReplicationFactor defines the desired replication factor; must be positive, or -1 to signify using the broker's default",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(-1),
						},
					},
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *KafkaBanzaicloudIoKafkaTopicV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_kafka_banzaicloud_io_kafka_topic_v1alpha1_manifest")

	var model KafkaBanzaicloudIoKafkaTopicV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("kafka.banzaicloud.io/v1alpha1")
	model.Kind = pointer.String("KafkaTopic")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
