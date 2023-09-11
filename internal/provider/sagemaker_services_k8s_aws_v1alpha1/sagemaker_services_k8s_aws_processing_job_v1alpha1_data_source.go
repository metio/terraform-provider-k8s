/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package sagemaker_services_k8s_aws_v1alpha1

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
	"net/http"
)

var (
	_ datasource.DataSource              = &SagemakerServicesK8SAwsProcessingJobV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &SagemakerServicesK8SAwsProcessingJobV1Alpha1DataSource{}
)

func NewSagemakerServicesK8SAwsProcessingJobV1Alpha1DataSource() datasource.DataSource {
	return &SagemakerServicesK8SAwsProcessingJobV1Alpha1DataSource{}
}

type SagemakerServicesK8SAwsProcessingJobV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type SagemakerServicesK8SAwsProcessingJobV1Alpha1DataSourceData struct {
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
		AppSpecification *struct {
			ContainerArguments  *[]string `tfsdk:"container_arguments" json:"containerArguments,omitempty"`
			ContainerEntrypoint *[]string `tfsdk:"container_entrypoint" json:"containerEntrypoint,omitempty"`
			ImageURI            *string   `tfsdk:"image_uri" json:"imageURI,omitempty"`
		} `tfsdk:"app_specification" json:"appSpecification,omitempty"`
		Environment      *map[string]string `tfsdk:"environment" json:"environment,omitempty"`
		ExperimentConfig *struct {
			ExperimentName            *string `tfsdk:"experiment_name" json:"experimentName,omitempty"`
			TrialComponentDisplayName *string `tfsdk:"trial_component_display_name" json:"trialComponentDisplayName,omitempty"`
			TrialName                 *string `tfsdk:"trial_name" json:"trialName,omitempty"`
		} `tfsdk:"experiment_config" json:"experimentConfig,omitempty"`
		NetworkConfig *struct {
			EnableInterContainerTrafficEncryption *bool `tfsdk:"enable_inter_container_traffic_encryption" json:"enableInterContainerTrafficEncryption,omitempty"`
			EnableNetworkIsolation                *bool `tfsdk:"enable_network_isolation" json:"enableNetworkIsolation,omitempty"`
			VpcConfig                             *struct {
				SecurityGroupIDs *[]string `tfsdk:"security_group_i_ds" json:"securityGroupIDs,omitempty"`
				Subnets          *[]string `tfsdk:"subnets" json:"subnets,omitempty"`
			} `tfsdk:"vpc_config" json:"vpcConfig,omitempty"`
		} `tfsdk:"network_config" json:"networkConfig,omitempty"`
		ProcessingInputs *[]struct {
			AppManaged        *bool `tfsdk:"app_managed" json:"appManaged,omitempty"`
			DatasetDefinition *struct {
				AthenaDatasetDefinition *struct {
					Catalog           *string `tfsdk:"catalog" json:"catalog,omitempty"`
					Database          *string `tfsdk:"database" json:"database,omitempty"`
					KmsKeyID          *string `tfsdk:"kms_key_id" json:"kmsKeyID,omitempty"`
					OutputCompression *string `tfsdk:"output_compression" json:"outputCompression,omitempty"`
					OutputFormat      *string `tfsdk:"output_format" json:"outputFormat,omitempty"`
					OutputS3URI       *string `tfsdk:"output_s3_uri" json:"outputS3URI,omitempty"`
					QueryString       *string `tfsdk:"query_string" json:"queryString,omitempty"`
					WorkGroup         *string `tfsdk:"work_group" json:"workGroup,omitempty"`
				} `tfsdk:"athena_dataset_definition" json:"athenaDatasetDefinition,omitempty"`
				DataDistributionType      *string `tfsdk:"data_distribution_type" json:"dataDistributionType,omitempty"`
				InputMode                 *string `tfsdk:"input_mode" json:"inputMode,omitempty"`
				LocalPath                 *string `tfsdk:"local_path" json:"localPath,omitempty"`
				RedshiftDatasetDefinition *struct {
					ClusterID         *string `tfsdk:"cluster_id" json:"clusterID,omitempty"`
					ClusterRoleARN    *string `tfsdk:"cluster_role_arn" json:"clusterRoleARN,omitempty"`
					Database          *string `tfsdk:"database" json:"database,omitempty"`
					DbUser            *string `tfsdk:"db_user" json:"dbUser,omitempty"`
					KmsKeyID          *string `tfsdk:"kms_key_id" json:"kmsKeyID,omitempty"`
					OutputCompression *string `tfsdk:"output_compression" json:"outputCompression,omitempty"`
					OutputFormat      *string `tfsdk:"output_format" json:"outputFormat,omitempty"`
					OutputS3URI       *string `tfsdk:"output_s3_uri" json:"outputS3URI,omitempty"`
					QueryString       *string `tfsdk:"query_string" json:"queryString,omitempty"`
				} `tfsdk:"redshift_dataset_definition" json:"redshiftDatasetDefinition,omitempty"`
			} `tfsdk:"dataset_definition" json:"datasetDefinition,omitempty"`
			InputName *string `tfsdk:"input_name" json:"inputName,omitempty"`
			S3Input   *struct {
				LocalPath              *string `tfsdk:"local_path" json:"localPath,omitempty"`
				S3CompressionType      *string `tfsdk:"s3_compression_type" json:"s3CompressionType,omitempty"`
				S3DataDistributionType *string `tfsdk:"s3_data_distribution_type" json:"s3DataDistributionType,omitempty"`
				S3DataType             *string `tfsdk:"s3_data_type" json:"s3DataType,omitempty"`
				S3InputMode            *string `tfsdk:"s3_input_mode" json:"s3InputMode,omitempty"`
				S3URI                  *string `tfsdk:"s3_uri" json:"s3URI,omitempty"`
			} `tfsdk:"s3_input" json:"s3Input,omitempty"`
		} `tfsdk:"processing_inputs" json:"processingInputs,omitempty"`
		ProcessingJobName      *string `tfsdk:"processing_job_name" json:"processingJobName,omitempty"`
		ProcessingOutputConfig *struct {
			KmsKeyID *string `tfsdk:"kms_key_id" json:"kmsKeyID,omitempty"`
			Outputs  *[]struct {
				AppManaged         *bool `tfsdk:"app_managed" json:"appManaged,omitempty"`
				FeatureStoreOutput *struct {
					FeatureGroupName *string `tfsdk:"feature_group_name" json:"featureGroupName,omitempty"`
				} `tfsdk:"feature_store_output" json:"featureStoreOutput,omitempty"`
				OutputName *string `tfsdk:"output_name" json:"outputName,omitempty"`
				S3Output   *struct {
					LocalPath    *string `tfsdk:"local_path" json:"localPath,omitempty"`
					S3URI        *string `tfsdk:"s3_uri" json:"s3URI,omitempty"`
					S3UploadMode *string `tfsdk:"s3_upload_mode" json:"s3UploadMode,omitempty"`
				} `tfsdk:"s3_output" json:"s3Output,omitempty"`
			} `tfsdk:"outputs" json:"outputs,omitempty"`
		} `tfsdk:"processing_output_config" json:"processingOutputConfig,omitempty"`
		ProcessingResources *struct {
			ClusterConfig *struct {
				InstanceCount  *int64  `tfsdk:"instance_count" json:"instanceCount,omitempty"`
				InstanceType   *string `tfsdk:"instance_type" json:"instanceType,omitempty"`
				VolumeKMSKeyID *string `tfsdk:"volume_kms_key_id" json:"volumeKMSKeyID,omitempty"`
				VolumeSizeInGB *int64  `tfsdk:"volume_size_in_gb" json:"volumeSizeInGB,omitempty"`
			} `tfsdk:"cluster_config" json:"clusterConfig,omitempty"`
		} `tfsdk:"processing_resources" json:"processingResources,omitempty"`
		RoleARN           *string `tfsdk:"role_arn" json:"roleARN,omitempty"`
		StoppingCondition *struct {
			MaxRuntimeInSeconds *int64 `tfsdk:"max_runtime_in_seconds" json:"maxRuntimeInSeconds,omitempty"`
		} `tfsdk:"stopping_condition" json:"stoppingCondition,omitempty"`
		Tags *[]struct {
			Key   *string `tfsdk:"key" json:"key,omitempty"`
			Value *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"tags" json:"tags,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *SagemakerServicesK8SAwsProcessingJobV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_sagemaker_services_k8s_aws_processing_job_v1alpha1"
}

func (r *SagemakerServicesK8SAwsProcessingJobV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ProcessingJob is the Schema for the ProcessingJobs API",
		MarkdownDescription: "ProcessingJob is the Schema for the ProcessingJobs API",
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
				Description:         "ProcessingJobSpec defines the desired state of ProcessingJob.  An Amazon SageMaker processing job that is used to analyze data and evaluate models. For more information, see Process Data and Evaluate Models (https://docs.aws.amazon.com/sagemaker/latest/dg/processing-job.html).",
				MarkdownDescription: "ProcessingJobSpec defines the desired state of ProcessingJob.  An Amazon SageMaker processing job that is used to analyze data and evaluate models. For more information, see Process Data and Evaluate Models (https://docs.aws.amazon.com/sagemaker/latest/dg/processing-job.html).",
				Attributes: map[string]schema.Attribute{
					"app_specification": schema.SingleNestedAttribute{
						Description:         "Configures the processing job to run a specified Docker container image.",
						MarkdownDescription: "Configures the processing job to run a specified Docker container image.",
						Attributes: map[string]schema.Attribute{
							"container_arguments": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"container_entrypoint": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"image_uri": schema.StringAttribute{
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

					"environment": schema.MapAttribute{
						Description:         "The environment variables to set in the Docker container. Up to 100 key and values entries in the map are supported.",
						MarkdownDescription: "The environment variables to set in the Docker container. Up to 100 key and values entries in the map are supported.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"experiment_config": schema.SingleNestedAttribute{
						Description:         "Associates a SageMaker job as a trial component with an experiment and trial. Specified when you call the following APIs:  * CreateProcessingJob  * CreateTrainingJob  * CreateTransformJob",
						MarkdownDescription: "Associates a SageMaker job as a trial component with an experiment and trial. Specified when you call the following APIs:  * CreateProcessingJob  * CreateTrainingJob  * CreateTransformJob",
						Attributes: map[string]schema.Attribute{
							"experiment_name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"trial_component_display_name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"trial_name": schema.StringAttribute{
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

					"network_config": schema.SingleNestedAttribute{
						Description:         "Networking options for a processing job, such as whether to allow inbound and outbound network calls to and from processing containers, and the VPC subnets and security groups to use for VPC-enabled processing jobs.",
						MarkdownDescription: "Networking options for a processing job, such as whether to allow inbound and outbound network calls to and from processing containers, and the VPC subnets and security groups to use for VPC-enabled processing jobs.",
						Attributes: map[string]schema.Attribute{
							"enable_inter_container_traffic_encryption": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"enable_network_isolation": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"vpc_config": schema.SingleNestedAttribute{
								Description:         "Specifies a VPC that your training jobs and hosted models have access to. Control access to and from your training and model containers by configuring the VPC. For more information, see Protect Endpoints by Using an Amazon Virtual Private Cloud (https://docs.aws.amazon.com/sagemaker/latest/dg/host-vpc.html) and Protect Training Jobs by Using an Amazon Virtual Private Cloud (https://docs.aws.amazon.com/sagemaker/latest/dg/train-vpc.html).",
								MarkdownDescription: "Specifies a VPC that your training jobs and hosted models have access to. Control access to and from your training and model containers by configuring the VPC. For more information, see Protect Endpoints by Using an Amazon Virtual Private Cloud (https://docs.aws.amazon.com/sagemaker/latest/dg/host-vpc.html) and Protect Training Jobs by Using an Amazon Virtual Private Cloud (https://docs.aws.amazon.com/sagemaker/latest/dg/train-vpc.html).",
								Attributes: map[string]schema.Attribute{
									"security_group_i_ds": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"subnets": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
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

					"processing_inputs": schema.ListNestedAttribute{
						Description:         "An array of inputs configuring the data to download into the processing container.",
						MarkdownDescription: "An array of inputs configuring the data to download into the processing container.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"app_managed": schema.BoolAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"dataset_definition": schema.SingleNestedAttribute{
									Description:         "Configuration for Dataset Definition inputs. The Dataset Definition input must specify exactly one of either AthenaDatasetDefinition or RedshiftDatasetDefinition types.",
									MarkdownDescription: "Configuration for Dataset Definition inputs. The Dataset Definition input must specify exactly one of either AthenaDatasetDefinition or RedshiftDatasetDefinition types.",
									Attributes: map[string]schema.Attribute{
										"athena_dataset_definition": schema.SingleNestedAttribute{
											Description:         "Configuration for Athena Dataset Definition input.",
											MarkdownDescription: "Configuration for Athena Dataset Definition input.",
											Attributes: map[string]schema.Attribute{
												"catalog": schema.StringAttribute{
													Description:         "The name of the data catalog used in Athena query execution.",
													MarkdownDescription: "The name of the data catalog used in Athena query execution.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"database": schema.StringAttribute{
													Description:         "The name of the database used in the Athena query execution.",
													MarkdownDescription: "The name of the database used in the Athena query execution.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"kms_key_id": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"output_compression": schema.StringAttribute{
													Description:         "The compression used for Athena query results.",
													MarkdownDescription: "The compression used for Athena query results.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"output_format": schema.StringAttribute{
													Description:         "The data storage format for Athena query results.",
													MarkdownDescription: "The data storage format for Athena query results.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"output_s3_uri": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"query_string": schema.StringAttribute{
													Description:         "The SQL query statements, to be executed.",
													MarkdownDescription: "The SQL query statements, to be executed.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"work_group": schema.StringAttribute{
													Description:         "The name of the workgroup in which the Athena query is being started.",
													MarkdownDescription: "The name of the workgroup in which the Athena query is being started.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"data_distribution_type": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"input_mode": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"local_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"redshift_dataset_definition": schema.SingleNestedAttribute{
											Description:         "Configuration for Redshift Dataset Definition input.",
											MarkdownDescription: "Configuration for Redshift Dataset Definition input.",
											Attributes: map[string]schema.Attribute{
												"cluster_id": schema.StringAttribute{
													Description:         "The Redshift cluster Identifier.",
													MarkdownDescription: "The Redshift cluster Identifier.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"cluster_role_arn": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"database": schema.StringAttribute{
													Description:         "The name of the Redshift database used in Redshift query execution.",
													MarkdownDescription: "The name of the Redshift database used in Redshift query execution.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"db_user": schema.StringAttribute{
													Description:         "The database user name used in Redshift query execution.",
													MarkdownDescription: "The database user name used in Redshift query execution.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"kms_key_id": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"output_compression": schema.StringAttribute{
													Description:         "The compression used for Redshift query results.",
													MarkdownDescription: "The compression used for Redshift query results.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"output_format": schema.StringAttribute{
													Description:         "The data storage format for Redshift query results.",
													MarkdownDescription: "The data storage format for Redshift query results.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"output_s3_uri": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"query_string": schema.StringAttribute{
													Description:         "The SQL query statements to be executed.",
													MarkdownDescription: "The SQL query statements to be executed.",
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

								"input_name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"s3_input": schema.SingleNestedAttribute{
									Description:         "Configuration for downloading input data from Amazon S3 into the processing container.",
									MarkdownDescription: "Configuration for downloading input data from Amazon S3 into the processing container.",
									Attributes: map[string]schema.Attribute{
										"local_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"s3_compression_type": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"s3_data_distribution_type": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"s3_data_type": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"s3_input_mode": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"s3_uri": schema.StringAttribute{
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
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"processing_job_name": schema.StringAttribute{
						Description:         "The name of the processing job. The name must be unique within an Amazon Web Services Region in the Amazon Web Services account.",
						MarkdownDescription: "The name of the processing job. The name must be unique within an Amazon Web Services Region in the Amazon Web Services account.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"processing_output_config": schema.SingleNestedAttribute{
						Description:         "Output configuration for the processing job.",
						MarkdownDescription: "Output configuration for the processing job.",
						Attributes: map[string]schema.Attribute{
							"kms_key_id": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"outputs": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"app_managed": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"feature_store_output": schema.SingleNestedAttribute{
											Description:         "Configuration for processing job outputs in Amazon SageMaker Feature Store.",
											MarkdownDescription: "Configuration for processing job outputs in Amazon SageMaker Feature Store.",
											Attributes: map[string]schema.Attribute{
												"feature_group_name": schema.StringAttribute{
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

										"output_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"s3_output": schema.SingleNestedAttribute{
											Description:         "Configuration for uploading output data to Amazon S3 from the processing container.",
											MarkdownDescription: "Configuration for uploading output data to Amazon S3 from the processing container.",
											Attributes: map[string]schema.Attribute{
												"local_path": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"s3_uri": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"s3_upload_mode": schema.StringAttribute{
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

					"processing_resources": schema.SingleNestedAttribute{
						Description:         "Identifies the resources, ML compute instances, and ML storage volumes to deploy for a processing job. In distributed training, you specify more than one instance.",
						MarkdownDescription: "Identifies the resources, ML compute instances, and ML storage volumes to deploy for a processing job. In distributed training, you specify more than one instance.",
						Attributes: map[string]schema.Attribute{
							"cluster_config": schema.SingleNestedAttribute{
								Description:         "Configuration for the cluster used to run a processing job.",
								MarkdownDescription: "Configuration for the cluster used to run a processing job.",
								Attributes: map[string]schema.Attribute{
									"instance_count": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"instance_type": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"volume_kms_key_id": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"volume_size_in_gb": schema.Int64Attribute{
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

					"role_arn": schema.StringAttribute{
						Description:         "The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume to perform tasks on your behalf.",
						MarkdownDescription: "The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume to perform tasks on your behalf.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"stopping_condition": schema.SingleNestedAttribute{
						Description:         "The time limit for how long the processing job is allowed to run.",
						MarkdownDescription: "The time limit for how long the processing job is allowed to run.",
						Attributes: map[string]schema.Attribute{
							"max_runtime_in_seconds": schema.Int64Attribute{
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

					"tags": schema.ListNestedAttribute{
						Description:         "(Optional) An array of key-value pairs. For more information, see Using Cost Allocation Tags (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html#allocation-whatURL) in the Amazon Web Services Billing and Cost Management User Guide.",
						MarkdownDescription: "(Optional) An array of key-value pairs. For more information, see Using Cost Allocation Tags (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html#allocation-whatURL) in the Amazon Web Services Billing and Cost Management User Guide.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"key": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"value": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},
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

func (r *SagemakerServicesK8SAwsProcessingJobV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *SagemakerServicesK8SAwsProcessingJobV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_sagemaker_services_k8s_aws_processing_job_v1alpha1")

	var data SagemakerServicesK8SAwsProcessingJobV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "sagemaker.services.k8s.aws", Version: "v1alpha1", Resource: "processingjobs"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		var statusError *k8sErrors.StatusError
		if errors.As(err, &statusError) {
			if statusError.Status().Code == http.StatusNotFound {
				response.Diagnostics.AddError(
					"Unable to find resource",
					fmt.Sprintf("The requested resource cannot be found. "+
						"Make sure that it does exist in your cluster and you have set the correct name and namespace configured.\n\n"+
						"Namespace: %s\n"+
						"Name: %s", data.Metadata.Namespace, data.Metadata.Name),
				)
				return
			}
		} else {
			response.Diagnostics.AddError(
				"Unable to GET resource",
				fmt.Sprintf("An unexpected error occurred while reading the resource. "+
					"Please report this issue to the provider developers.\n\n"+
					"GET Error (%T): %s", err, err.Error()),
			)
		}
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

	var readResponse SagemakerServicesK8SAwsProcessingJobV1Alpha1DataSourceData
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
	data.ApiVersion = pointer.String("sagemaker.services.k8s.aws/v1alpha1")
	data.Kind = pointer.String("ProcessingJob")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
