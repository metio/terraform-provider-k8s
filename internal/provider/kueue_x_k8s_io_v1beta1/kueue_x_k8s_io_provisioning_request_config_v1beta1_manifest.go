/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package kueue_x_k8s_io_v1beta1

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
	"regexp"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &KueueXK8SIoProvisioningRequestConfigV1Beta1Manifest{}
)

func NewKueueXK8SIoProvisioningRequestConfigV1Beta1Manifest() datasource.DataSource {
	return &KueueXK8SIoProvisioningRequestConfigV1Beta1Manifest{}
}

type KueueXK8SIoProvisioningRequestConfigV1Beta1Manifest struct{}

type KueueXK8SIoProvisioningRequestConfigV1Beta1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		ManagedResources *[]string          `tfsdk:"managed_resources" json:"managedResources,omitempty"`
		Parameters       *map[string]string `tfsdk:"parameters" json:"parameters,omitempty"`
		PodSetUpdates    *struct {
			NodeSelector *[]struct {
				Key                              *string `tfsdk:"key" json:"key,omitempty"`
				ValueFromProvisioningClassDetail *string `tfsdk:"value_from_provisioning_class_detail" json:"valueFromProvisioningClassDetail,omitempty"`
			} `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
		} `tfsdk:"pod_set_updates" json:"podSetUpdates,omitempty"`
		ProvisioningClassName *string `tfsdk:"provisioning_class_name" json:"provisioningClassName,omitempty"`
		RetryStrategy         *struct {
			BackoffBaseSeconds *int64 `tfsdk:"backoff_base_seconds" json:"backoffBaseSeconds,omitempty"`
			BackoffLimitCount  *int64 `tfsdk:"backoff_limit_count" json:"backoffLimitCount,omitempty"`
			BackoffMaxSeconds  *int64 `tfsdk:"backoff_max_seconds" json:"backoffMaxSeconds,omitempty"`
		} `tfsdk:"retry_strategy" json:"retryStrategy,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *KueueXK8SIoProvisioningRequestConfigV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_kueue_x_k8s_io_provisioning_request_config_v1beta1_manifest"
}

func (r *KueueXK8SIoProvisioningRequestConfigV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ProvisioningRequestConfig is the Schema for the provisioningrequestconfig API",
		MarkdownDescription: "ProvisioningRequestConfig is the Schema for the provisioningrequestconfig API",
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
				Description:         "ProvisioningRequestConfigSpec defines the desired state of ProvisioningRequestConfig",
				MarkdownDescription: "ProvisioningRequestConfigSpec defines the desired state of ProvisioningRequestConfig",
				Attributes: map[string]schema.Attribute{
					"managed_resources": schema.ListAttribute{
						Description:         "managedResources contains the list of resources managed by the autoscaling. If empty, all resources are considered managed. If not empty, the ProvisioningRequest will contain only the podsets that are requesting at least one of them. If none of the workloads podsets is requesting at least a managed resource, the workload is considered ready.",
						MarkdownDescription: "managedResources contains the list of resources managed by the autoscaling. If empty, all resources are considered managed. If not empty, the ProvisioningRequest will contain only the podsets that are requesting at least one of them. If none of the workloads podsets is requesting at least a managed resource, the workload is considered ready.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"parameters": schema.MapAttribute{
						Description:         "Parameters contains all other parameters classes may require.",
						MarkdownDescription: "Parameters contains all other parameters classes may require.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"pod_set_updates": schema.SingleNestedAttribute{
						Description:         "podSetUpdates specifies the update of the workload's PodSetUpdates which are used to target the provisioned nodes.",
						MarkdownDescription: "podSetUpdates specifies the update of the workload's PodSetUpdates which are used to target the provisioned nodes.",
						Attributes: map[string]schema.Attribute{
							"node_selector": schema.ListNestedAttribute{
								Description:         "nodeSelector specifies the list of updates for the NodeSelector.",
								MarkdownDescription: "nodeSelector specifies the list of updates for the NodeSelector.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"key": schema.StringAttribute{
											Description:         "key specifies the key for the NodeSelector.",
											MarkdownDescription: "key specifies the key for the NodeSelector.",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtLeast(1),
												stringvalidator.LengthAtMost(317),
												stringvalidator.RegexMatches(regexp.MustCompile(`^([a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*/)?(([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9])$`), ""),
											},
										},

										"value_from_provisioning_class_detail": schema.StringAttribute{
											Description:         "valueFromProvisioningClassDetail specifies the key of the ProvisioningRequest.status.provisioningClassDetails from which the value is used for the update.",
											MarkdownDescription: "valueFromProvisioningClassDetail specifies the key of the ProvisioningRequest.status.provisioningClassDetails from which the value is used for the update.",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtLeast(1),
												stringvalidator.LengthAtMost(32768),
											},
										},
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

					"provisioning_class_name": schema.StringAttribute{
						Description:         "ProvisioningClassName describes the different modes of provisioning the resources. Check autoscaling.x-k8s.io ProvisioningRequestSpec.ProvisioningClassName for details.",
						MarkdownDescription: "ProvisioningClassName describes the different modes of provisioning the resources. Check autoscaling.x-k8s.io ProvisioningRequestSpec.ProvisioningClassName for details.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(253),
							stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
						},
					},

					"retry_strategy": schema.SingleNestedAttribute{
						Description:         "retryStrategy defines strategy for retrying ProvisioningRequest. If null, then the default configuration is applied with the following parameter values: backoffLimitCount: 3 backoffBaseSeconds: 60 - 1 min backoffMaxSeconds: 1800 - 30 mins To switch off retry mechanism set retryStrategy.backoffLimitCount to 0.",
						MarkdownDescription: "retryStrategy defines strategy for retrying ProvisioningRequest. If null, then the default configuration is applied with the following parameter values: backoffLimitCount: 3 backoffBaseSeconds: 60 - 1 min backoffMaxSeconds: 1800 - 30 mins To switch off retry mechanism set retryStrategy.backoffLimitCount to 0.",
						Attributes: map[string]schema.Attribute{
							"backoff_base_seconds": schema.Int64Attribute{
								Description:         "BackoffBaseSeconds defines the base for the exponential backoff for re-queuing an evicted workload. Defaults to 60.",
								MarkdownDescription: "BackoffBaseSeconds defines the base for the exponential backoff for re-queuing an evicted workload. Defaults to 60.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"backoff_limit_count": schema.Int64Attribute{
								Description:         "BackoffLimitCount defines the maximum number of re-queuing retries. Once the number is reached, the workload is deactivated ('.spec.activate'='false'). Every backoff duration is about 'b*2^(n-1)+Rand' where: - 'b' represents the base set by 'BackoffBaseSeconds' parameter, - 'n' represents the 'workloadStatus.requeueState.count', - 'Rand' represents the random jitter. During this time, the workload is taken as an inadmissible and other workloads will have a chance to be admitted. By default, the consecutive requeue delays are around: (60s, 120s, 240s, ...). Defaults to 3.",
								MarkdownDescription: "BackoffLimitCount defines the maximum number of re-queuing retries. Once the number is reached, the workload is deactivated ('.spec.activate'='false'). Every backoff duration is about 'b*2^(n-1)+Rand' where: - 'b' represents the base set by 'BackoffBaseSeconds' parameter, - 'n' represents the 'workloadStatus.requeueState.count', - 'Rand' represents the random jitter. During this time, the workload is taken as an inadmissible and other workloads will have a chance to be admitted. By default, the consecutive requeue delays are around: (60s, 120s, 240s, ...). Defaults to 3.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"backoff_max_seconds": schema.Int64Attribute{
								Description:         "BackoffMaxSeconds defines the maximum backoff time to re-queue an evicted workload. Defaults to 1800.",
								MarkdownDescription: "BackoffMaxSeconds defines the maximum backoff time to re-queue an evicted workload. Defaults to 1800.",
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

func (r *KueueXK8SIoProvisioningRequestConfigV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_kueue_x_k8s_io_provisioning_request_config_v1beta1_manifest")

	var model KueueXK8SIoProvisioningRequestConfigV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("kueue.x-k8s.io/v1beta1")
	model.Kind = pointer.String("ProvisioningRequestConfig")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
