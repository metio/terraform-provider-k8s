/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package emrcontainers_services_k8s_aws_v1alpha1

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
	_ datasource.DataSource              = &EmrcontainersServicesK8SAwsJobRunV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &EmrcontainersServicesK8SAwsJobRunV1Alpha1DataSource{}
)

func NewEmrcontainersServicesK8SAwsJobRunV1Alpha1DataSource() datasource.DataSource {
	return &EmrcontainersServicesK8SAwsJobRunV1Alpha1DataSource{}
}

type EmrcontainersServicesK8SAwsJobRunV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type EmrcontainersServicesK8SAwsJobRunV1Alpha1DataSourceData struct {
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
		ConfigurationOverrides *string `tfsdk:"configuration_overrides" json:"configurationOverrides,omitempty"`
		ExecutionRoleARN       *string `tfsdk:"execution_role_arn" json:"executionRoleARN,omitempty"`
		JobDriver              *struct {
			SparkSubmitJobDriver *struct {
				EntryPoint            *string   `tfsdk:"entry_point" json:"entryPoint,omitempty"`
				EntryPointArguments   *[]string `tfsdk:"entry_point_arguments" json:"entryPointArguments,omitempty"`
				SparkSubmitParameters *string   `tfsdk:"spark_submit_parameters" json:"sparkSubmitParameters,omitempty"`
			} `tfsdk:"spark_submit_job_driver" json:"sparkSubmitJobDriver,omitempty"`
		} `tfsdk:"job_driver" json:"jobDriver,omitempty"`
		Name              *string            `tfsdk:"name" json:"name,omitempty"`
		ReleaseLabel      *string            `tfsdk:"release_label" json:"releaseLabel,omitempty"`
		Tags              *map[string]string `tfsdk:"tags" json:"tags,omitempty"`
		VirtualClusterID  *string            `tfsdk:"virtual_cluster_id" json:"virtualClusterID,omitempty"`
		VirtualClusterRef *struct {
			From *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"from" json:"from,omitempty"`
		} `tfsdk:"virtual_cluster_ref" json:"virtualClusterRef,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *EmrcontainersServicesK8SAwsJobRunV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_emrcontainers_services_k8s_aws_job_run_v1alpha1"
}

func (r *EmrcontainersServicesK8SAwsJobRunV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "JobRun is the Schema for the JobRuns API",
		MarkdownDescription: "JobRun is the Schema for the JobRuns API",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"api_version": schema.StringAttribute{
				Description:         "The API group of the requested resource.",
				MarkdownDescription: "The API group of the requested resource.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"kind": schema.StringAttribute{
				Description:         "The type of the requested resource.",
				MarkdownDescription: "The type of the requested resource.",
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
				Description:         "JobRunSpec defines the desired state of JobRun.  This entity describes a job run. A job run is a unit of work, such as a Spark jar, PySpark script, or SparkSQL query, that you submit to Amazon EMR on EKS.",
				MarkdownDescription: "JobRunSpec defines the desired state of JobRun.  This entity describes a job run. A job run is a unit of work, such as a Spark jar, PySpark script, or SparkSQL query, that you submit to Amazon EMR on EKS.",
				Attributes: map[string]schema.Attribute{
					"configuration_overrides": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"execution_role_arn": schema.StringAttribute{
						Description:         "The execution role ARN for the job run.",
						MarkdownDescription: "The execution role ARN for the job run.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"job_driver": schema.SingleNestedAttribute{
						Description:         "The job driver for the job run.",
						MarkdownDescription: "The job driver for the job run.",
						Attributes: map[string]schema.Attribute{
							"spark_submit_job_driver": schema.SingleNestedAttribute{
								Description:         "The information about job driver for Spark submit.",
								MarkdownDescription: "The information about job driver for Spark submit.",
								Attributes: map[string]schema.Attribute{
									"entry_point": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"entry_point_arguments": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"spark_submit_parameters": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
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
						Required: false,
						Optional: false,
						Computed: true,
					},

					"name": schema.StringAttribute{
						Description:         "The name of the job run.",
						MarkdownDescription: "The name of the job run.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"release_label": schema.StringAttribute{
						Description:         "The Amazon EMR release version to use for the job run.",
						MarkdownDescription: "The Amazon EMR release version to use for the job run.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"tags": schema.MapAttribute{
						Description:         "The tags assigned to job runs.",
						MarkdownDescription: "The tags assigned to job runs.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"virtual_cluster_id": schema.StringAttribute{
						Description:         "The virtual cluster ID for which the job run request is submitted.",
						MarkdownDescription: "The virtual cluster ID for which the job run request is submitted.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"virtual_cluster_ref": schema.SingleNestedAttribute{
						Description:         "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReference type to provide more user friendly syntax for references using 'from' field Ex: APIIDRef:  from: name: my-api",
						MarkdownDescription: "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReference type to provide more user friendly syntax for references using 'from' field Ex: APIIDRef:  from: name: my-api",
						Attributes: map[string]schema.Attribute{
							"from": schema.SingleNestedAttribute{
								Description:         "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",
								MarkdownDescription: "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
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
						Required: false,
						Optional: false,
						Computed: true,
					},
				},
				Required: false,
				Optional: false,
				Computed: true,
			},
		},
	}
}

func (r *EmrcontainersServicesK8SAwsJobRunV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.Append(utilities.OfflineProviderError())
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.Append(utilities.UnexpectedDataSourceDataError(request.ProviderData))
	}
}

func (r *EmrcontainersServicesK8SAwsJobRunV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_emrcontainers_services_k8s_aws_job_run_v1alpha1")

	var data EmrcontainersServicesK8SAwsJobRunV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "emrcontainers.services.k8s.aws", Version: "v1alpha1", Resource: "jobruns"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.Append(utilities.GetNamespacedResourceError(err, data.Metadata.Name, data.Metadata.Namespace))
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse EmrcontainersServicesK8SAwsJobRunV1Alpha1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Namespace, data.Metadata.Name))
	data.ApiVersion = pointer.String("emrcontainers.services.k8s.aws/v1alpha1")
	data.Kind = pointer.String("JobRun")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
