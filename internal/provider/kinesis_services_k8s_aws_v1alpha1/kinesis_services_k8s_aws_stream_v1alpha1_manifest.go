/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package kinesis_services_k8s_aws_v1alpha1

import (
	"context"
	"fmt"
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
	_ datasource.DataSource = &KinesisServicesK8SAwsStreamV1Alpha1Manifest{}
)

func NewKinesisServicesK8SAwsStreamV1Alpha1Manifest() datasource.DataSource {
	return &KinesisServicesK8SAwsStreamV1Alpha1Manifest{}
}

type KinesisServicesK8SAwsStreamV1Alpha1Manifest struct{}

type KinesisServicesK8SAwsStreamV1Alpha1ManifestData struct {
	ID   types.String `tfsdk:"id" json:"-"`
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
		Name              *string `tfsdk:"name" json:"name,omitempty"`
		ShardCount        *int64  `tfsdk:"shard_count" json:"shardCount,omitempty"`
		StreamModeDetails *struct {
			StreamMode *string `tfsdk:"stream_mode" json:"streamMode,omitempty"`
		} `tfsdk:"stream_mode_details" json:"streamModeDetails,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *KinesisServicesK8SAwsStreamV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_kinesis_services_k8s_aws_stream_v1alpha1_manifest"
}

func (r *KinesisServicesK8SAwsStreamV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Stream is the Schema for the Streams API",
		MarkdownDescription: "Stream is the Schema for the Streams API",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

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
				Description:         "StreamSpec defines the desired state of Stream.",
				MarkdownDescription: "StreamSpec defines the desired state of Stream.",
				Attributes: map[string]schema.Attribute{
					"name": schema.StringAttribute{
						Description:         "A name to identify the stream. The stream name is scoped to the Amazon WebServices account used by the application that creates the stream. It is alsoscoped by Amazon Web Services Region. That is, two streams in two differentAmazon Web Services accounts can have the same name. Two streams in the sameAmazon Web Services account but in two different Regions can also have thesame name.",
						MarkdownDescription: "A name to identify the stream. The stream name is scoped to the Amazon WebServices account used by the application that creates the stream. It is alsoscoped by Amazon Web Services Region. That is, two streams in two differentAmazon Web Services accounts can have the same name. Two streams in the sameAmazon Web Services account but in two different Regions can also have thesame name.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"shard_count": schema.Int64Attribute{
						Description:         "The number of shards that the stream will use. The throughput of the streamis a function of the number of shards; more shards are required for greaterprovisioned throughput.",
						MarkdownDescription: "The number of shards that the stream will use. The throughput of the streamis a function of the number of shards; more shards are required for greaterprovisioned throughput.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"stream_mode_details": schema.SingleNestedAttribute{
						Description:         "Indicates the capacity mode of the data stream. Currently, in Kinesis DataStreams, you can choose between an on-demand capacity mode and a provisionedcapacity mode for your data streams.",
						MarkdownDescription: "Indicates the capacity mode of the data stream. Currently, in Kinesis DataStreams, you can choose between an on-demand capacity mode and a provisionedcapacity mode for your data streams.",
						Attributes: map[string]schema.Attribute{
							"stream_mode": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
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
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *KinesisServicesK8SAwsStreamV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_kinesis_services_k8s_aws_stream_v1alpha1_manifest")

	var model KinesisServicesK8SAwsStreamV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("kinesis.services.k8s.aws/v1alpha1")
	model.Kind = pointer.String("Stream")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
