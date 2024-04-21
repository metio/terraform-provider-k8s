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
	_ datasource.DataSource = &KafkaBanzaicloudIoCruiseControlOperationV1Alpha1Manifest{}
)

func NewKafkaBanzaicloudIoCruiseControlOperationV1Alpha1Manifest() datasource.DataSource {
	return &KafkaBanzaicloudIoCruiseControlOperationV1Alpha1Manifest{}
}

type KafkaBanzaicloudIoCruiseControlOperationV1Alpha1Manifest struct{}

type KafkaBanzaicloudIoCruiseControlOperationV1Alpha1ManifestData struct {
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
		ErrorPolicy             *string `tfsdk:"error_policy" json:"errorPolicy,omitempty"`
		TtlSecondsAfterFinished *int64  `tfsdk:"ttl_seconds_after_finished" json:"ttlSecondsAfterFinished,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *KafkaBanzaicloudIoCruiseControlOperationV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_kafka_banzaicloud_io_cruise_control_operation_v1alpha1_manifest"
}

func (r *KafkaBanzaicloudIoCruiseControlOperationV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "CruiseControlOperation is the Schema for the cruiseControlOperation API.",
		MarkdownDescription: "CruiseControlOperation is the Schema for the cruiseControlOperation API.",
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
				Description:         "CruiseControlOperationSpec defines the desired state of CruiseControlOperation.",
				MarkdownDescription: "CruiseControlOperationSpec defines the desired state of CruiseControlOperation.",
				Attributes: map[string]schema.Attribute{
					"error_policy": schema.StringAttribute{
						Description:         "ErrorPolicy defines how failed Cruise Control operation should be handled. When it is 'retry', the Koperator re-executes the failed task in every 30 sec (by default). When it is 'ignore', the Koperator handles the failed task as completed.",
						MarkdownDescription: "ErrorPolicy defines how failed Cruise Control operation should be handled. When it is 'retry', the Koperator re-executes the failed task in every 30 sec (by default). When it is 'ignore', the Koperator handles the failed task as completed.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("ignore", "retry"),
						},
					},

					"ttl_seconds_after_finished": schema.Int64Attribute{
						Description:         "When TTLSecondsAfterFinished is specified, the created and finished (completed successfully or completedWithError and errorPolicy: ignore) cruiseControlOperation custom resource will be deleted after the given time elapsed. When it is 0 then the resource is going to be deleted instantly after the operation is finished. When it is not specified the resource is not going to be removed. Value can be only zero and positive integers",
						MarkdownDescription: "When TTLSecondsAfterFinished is specified, the created and finished (completed successfully or completedWithError and errorPolicy: ignore) cruiseControlOperation custom resource will be deleted after the given time elapsed. When it is 0 then the resource is going to be deleted instantly after the operation is finished. When it is not specified the resource is not going to be removed. Value can be only zero and positive integers",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
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

func (r *KafkaBanzaicloudIoCruiseControlOperationV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_kafka_banzaicloud_io_cruise_control_operation_v1alpha1_manifest")

	var model KafkaBanzaicloudIoCruiseControlOperationV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("kafka.banzaicloud.io/v1alpha1")
	model.Kind = pointer.String("CruiseControlOperation")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
