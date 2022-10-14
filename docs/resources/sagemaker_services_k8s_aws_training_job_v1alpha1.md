---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_sagemaker_services_k8s_aws_training_job_v1alpha1 Resource - terraform-provider-k8s"
subcategory: "sagemaker.services.k8s.aws/v1alpha1"
description: |-
  TrainingJob is the Schema for the TrainingJobs API
---

# k8s_sagemaker_services_k8s_aws_training_job_v1alpha1 (Resource)

TrainingJob is the Schema for the TrainingJobs API

## Example Usage

```terraform
resource "k8s_sagemaker_services_k8s_aws_training_job_v1alpha1" "minimal" {
  metadata = {
    name = "test"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `metadata` (Attributes) Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details. (see [below for nested schema](#nestedatt--metadata))

### Optional

- `spec` (Attributes) TrainingJobSpec defines the desired state of TrainingJob.  Contains information about a training job. (see [below for nested schema](#nestedatt--spec))

### Read-Only

- `api_version` (String) APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
- `id` (Number) The timestamp of the last change to this resource.
- `kind` (String) Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
- `yaml` (String) The generated manifest in YAML format.

<a id="nestedatt--metadata"></a>
### Nested Schema for `metadata`

Required:

- `name` (String) Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.

Optional:

- `annotations` (Map of String) Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.
- `labels` (Map of String) Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.
- `namespace` (String) Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.


<a id="nestedatt--spec"></a>
### Nested Schema for `spec`

Required:

- `algorithm_specification` (Attributes) The registry path of the Docker image that contains the training algorithm and algorithm-specific metadata, including the input mode. For more information about algorithms provided by SageMaker, see Algorithms (https://docs.aws.amazon.com/sagemaker/latest/dg/algos.html). For information about providing your own algorithms, see Using Your Own Algorithms with Amazon SageMaker (https://docs.aws.amazon.com/sagemaker/latest/dg/your-algorithms.html). (see [below for nested schema](#nestedatt--spec--algorithm_specification))
- `output_data_config` (Attributes) Specifies the path to the S3 location where you want to store model artifacts. SageMaker creates subfolders for the artifacts. (see [below for nested schema](#nestedatt--spec--output_data_config))
- `resource_config` (Attributes) The resources, including the ML compute instances and ML storage volumes, to use for model training.  ML storage volumes store model artifacts and incremental states. Training algorithms might also use ML storage volumes for scratch space. If you want SageMaker to use the ML storage volume to store the training data, choose File as the TrainingInputMode in the algorithm specification. For distributed training algorithms, specify an instance count greater than 1. (see [below for nested schema](#nestedatt--spec--resource_config))
- `role_arn` (String) The Amazon Resource Name (ARN) of an IAM role that SageMaker can assume to perform tasks on your behalf.  During model training, SageMaker needs your permission to read input data from an S3 bucket, download a Docker image that contains training code, write model artifacts to an S3 bucket, write logs to Amazon CloudWatch Logs, and publish metrics to Amazon CloudWatch. You grant permissions for all of these tasks to an IAM role. For more information, see SageMaker Roles (https://docs.aws.amazon.com/sagemaker/latest/dg/sagemaker-roles.html).  To be able to pass this role to SageMaker, the caller of this API must have the iam:PassRole permission.
- `stopping_condition` (Attributes) Specifies a limit to how long a model training job can run. It also specifies how long a managed Spot training job has to complete. When the job reaches the time limit, SageMaker ends the training job. Use this API to cap model training costs.  To stop a job, SageMaker sends the algorithm the SIGTERM signal, which delays job termination for 120 seconds. Algorithms can use this 120-second window to save the model artifacts, so the results of training are not lost. (see [below for nested schema](#nestedatt--spec--stopping_condition))
- `training_job_name` (String) The name of the training job. The name must be unique within an Amazon Web Services Region in an Amazon Web Services account.

Optional:

- `checkpoint_config` (Attributes) Contains information about the output location for managed spot training checkpoint data. (see [below for nested schema](#nestedatt--spec--checkpoint_config))
- `debug_hook_config` (Attributes) Configuration information for the Debugger hook parameters, metric and tensor collections, and storage paths. To learn more about how to configure the DebugHookConfig parameter, see Use the SageMaker and Debugger Configuration API Operations to Create, Update, and Debug Your Training Job (https://docs.aws.amazon.com/sagemaker/latest/dg/debugger-createtrainingjob-api.html). (see [below for nested schema](#nestedatt--spec--debug_hook_config))
- `debug_rule_configurations` (Attributes List) Configuration information for Debugger rules for debugging output tensors. (see [below for nested schema](#nestedatt--spec--debug_rule_configurations))
- `enable_inter_container_traffic_encryption` (Boolean) To encrypt all communications between ML compute instances in distributed training, choose True. Encryption provides greater security for distributed training, but training might take longer. How long it takes depends on the amount of communication between compute instances, especially if you use a deep learning algorithm in distributed training. For more information, see Protect Communications Between ML Compute Instances in a Distributed Training Job (https://docs.aws.amazon.com/sagemaker/latest/dg/train-encrypt.html).
- `enable_managed_spot_training` (Boolean) To train models using managed spot training, choose True. Managed spot training provides a fully managed and scalable infrastructure for training machine learning models. this option is useful when training jobs can be interrupted and when there is flexibility when the training job is run.  The complete and intermediate results of jobs are stored in an Amazon S3 bucket, and can be used as a starting point to train models incrementally. Amazon SageMaker provides metrics and logs in CloudWatch. They can be used to see when managed spot training jobs are running, interrupted, resumed, or completed.
- `enable_network_isolation` (Boolean) Isolates the training container. No inbound or outbound network calls can be made, except for calls between peers within a training cluster for distributed training. If you enable network isolation for training jobs that are configured to use a VPC, SageMaker downloads and uploads customer data and model artifacts through the specified VPC, but the training container does not have network access.
- `environment` (Map of String) The environment variables to set in the Docker container.
- `experiment_config` (Attributes) Associates a SageMaker job as a trial component with an experiment and trial. Specified when you call the following APIs:  * CreateProcessingJob  * CreateTrainingJob  * CreateTransformJob (see [below for nested schema](#nestedatt--spec--experiment_config))
- `hyper_parameters` (Map of String) Algorithm-specific parameters that influence the quality of the model. You set hyperparameters before you start the learning process. For a list of hyperparameters for each training algorithm provided by SageMaker, see Algorithms (https://docs.aws.amazon.com/sagemaker/latest/dg/algos.html).  You can specify a maximum of 100 hyperparameters. Each hyperparameter is a key-value pair. Each key and value is limited to 256 characters, as specified by the Length Constraint.
- `input_data_config` (Attributes List) An array of Channel objects. Each channel is a named input source. InputDataConfig describes the input data and its location.  Algorithms can accept input data from one or more channels. For example, an algorithm might have two channels of input data, training_data and validation_data. The configuration for each channel provides the S3, EFS, or FSx location where the input data is stored. It also provides information about the stored data: the MIME type, compression method, and whether the data is wrapped in RecordIO format.  Depending on the input mode that the algorithm supports, SageMaker either copies input data files from an S3 bucket to a local directory in the Docker container, or makes it available as input streams. For example, if you specify an EFS location, input data files are available as input streams. They do not need to be downloaded. (see [below for nested schema](#nestedatt--spec--input_data_config))
- `profiler_config` (Attributes) Configuration information for Debugger system monitoring, framework profiling, and storage paths. (see [below for nested schema](#nestedatt--spec--profiler_config))
- `profiler_rule_configurations` (Attributes List) Configuration information for Debugger rules for profiling system and framework metrics. (see [below for nested schema](#nestedatt--spec--profiler_rule_configurations))
- `tags` (Attributes List) An array of key-value pairs. You can use tags to categorize your Amazon Web Services resources in different ways, for example, by purpose, owner, or environment. For more information, see Tagging Amazon Web Services Resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html). (see [below for nested schema](#nestedatt--spec--tags))
- `tensor_board_output_config` (Attributes) Configuration of storage locations for the Debugger TensorBoard output data. (see [below for nested schema](#nestedatt--spec--tensor_board_output_config))
- `vpc_config` (Attributes) A VpcConfig object that specifies the VPC that you want your training job to connect to. Control access to and from your training container by configuring the VPC. For more information, see Protect Training Jobs by Using an Amazon Virtual Private Cloud (https://docs.aws.amazon.com/sagemaker/latest/dg/train-vpc.html). (see [below for nested schema](#nestedatt--spec--vpc_config))

<a id="nestedatt--spec--algorithm_specification"></a>
### Nested Schema for `spec.algorithm_specification`

Optional:

- `algorithm_name` (String)
- `enable_sage_maker_metrics_time_series` (Boolean)
- `metric_definitions` (Attributes List) (see [below for nested schema](#nestedatt--spec--algorithm_specification--metric_definitions))
- `training_image` (String)
- `training_input_mode` (String) The training input mode that the algorithm supports. For more information about input modes, see Algorithms (https://docs.aws.amazon.com/sagemaker/latest/dg/algos.html).  Pipe mode  If an algorithm supports Pipe mode, Amazon SageMaker streams data directly from Amazon S3 to the container.  File mode  If an algorithm supports File mode, SageMaker downloads the training data from S3 to the provisioned ML storage volume, and mounts the directory to the Docker volume for the training container.  You must provision the ML storage volume with sufficient capacity to accommodate the data downloaded from S3. In addition to the training data, the ML storage volume also stores the output model. The algorithm container uses the ML storage volume to also store intermediate information, if any.  For distributed algorithms, training data is distributed uniformly. Your training duration is predictable if the input data objects sizes are approximately the same. SageMaker does not split the files any further for model training. If the object sizes are skewed, training won't be optimal as the data distribution is also skewed when one host in a training cluster is overloaded, thus becoming a bottleneck in training.  FastFile mode  If an algorithm supports FastFile mode, SageMaker streams data directly from S3 to the container with no code changes, and provides file system access to the data. Users can author their training script to interact with these files as if they were stored on disk.  FastFile mode works best when the data is read sequentially. Augmented manifest files aren't supported. The startup time is lower when there are fewer files in the S3 bucket provided.

<a id="nestedatt--spec--algorithm_specification--metric_definitions"></a>
### Nested Schema for `spec.algorithm_specification.metric_definitions`

Optional:

- `name` (String)
- `regex` (String)



<a id="nestedatt--spec--output_data_config"></a>
### Nested Schema for `spec.output_data_config`

Optional:

- `kms_key_id` (String)
- `s3_output_path` (String)


<a id="nestedatt--spec--resource_config"></a>
### Nested Schema for `spec.resource_config`

Optional:

- `instance_count` (Number)
- `instance_type` (String)
- `volume_kms_key_id` (String)
- `volume_size_in_gb` (Number)


<a id="nestedatt--spec--stopping_condition"></a>
### Nested Schema for `spec.stopping_condition`

Optional:

- `max_runtime_in_seconds` (Number)
- `max_wait_time_in_seconds` (Number)


<a id="nestedatt--spec--checkpoint_config"></a>
### Nested Schema for `spec.checkpoint_config`

Optional:

- `local_path` (String)
- `s3_uri` (String)


<a id="nestedatt--spec--debug_hook_config"></a>
### Nested Schema for `spec.debug_hook_config`

Optional:

- `collection_configurations` (Attributes List) (see [below for nested schema](#nestedatt--spec--debug_hook_config--collection_configurations))
- `hook_parameters` (Map of String)
- `local_path` (String)
- `s3_output_path` (String)

<a id="nestedatt--spec--debug_hook_config--collection_configurations"></a>
### Nested Schema for `spec.debug_hook_config.collection_configurations`

Optional:

- `collection_name` (String)
- `collection_parameters` (Map of String)



<a id="nestedatt--spec--debug_rule_configurations"></a>
### Nested Schema for `spec.debug_rule_configurations`

Optional:

- `instance_type` (String)
- `local_path` (String)
- `rule_configuration_name` (String)
- `rule_evaluator_image` (String)
- `rule_parameters` (Map of String)
- `s3_output_path` (String)
- `volume_size_in_gb` (Number)


<a id="nestedatt--spec--experiment_config"></a>
### Nested Schema for `spec.experiment_config`

Optional:

- `experiment_name` (String)
- `trial_component_display_name` (String)
- `trial_name` (String)


<a id="nestedatt--spec--input_data_config"></a>
### Nested Schema for `spec.input_data_config`

Optional:

- `channel_name` (String)
- `compression_type` (String)
- `content_type` (String)
- `data_source` (Attributes) Describes the location of the channel data. (see [below for nested schema](#nestedatt--spec--input_data_config--data_source))
- `input_mode` (String) The training input mode that the algorithm supports. For more information about input modes, see Algorithms (https://docs.aws.amazon.com/sagemaker/latest/dg/algos.html).  Pipe mode  If an algorithm supports Pipe mode, Amazon SageMaker streams data directly from Amazon S3 to the container.  File mode  If an algorithm supports File mode, SageMaker downloads the training data from S3 to the provisioned ML storage volume, and mounts the directory to the Docker volume for the training container.  You must provision the ML storage volume with sufficient capacity to accommodate the data downloaded from S3. In addition to the training data, the ML storage volume also stores the output model. The algorithm container uses the ML storage volume to also store intermediate information, if any.  For distributed algorithms, training data is distributed uniformly. Your training duration is predictable if the input data objects sizes are approximately the same. SageMaker does not split the files any further for model training. If the object sizes are skewed, training won't be optimal as the data distribution is also skewed when one host in a training cluster is overloaded, thus becoming a bottleneck in training.  FastFile mode  If an algorithm supports FastFile mode, SageMaker streams data directly from S3 to the container with no code changes, and provides file system access to the data. Users can author their training script to interact with these files as if they were stored on disk.  FastFile mode works best when the data is read sequentially. Augmented manifest files aren't supported. The startup time is lower when there are fewer files in the S3 bucket provided.
- `record_wrapper_type` (String)
- `shuffle_config` (Attributes) A configuration for a shuffle option for input data in a channel. If you use S3Prefix for S3DataType, the results of the S3 key prefix matches are shuffled. If you use ManifestFile, the order of the S3 object references in the ManifestFile is shuffled. If you use AugmentedManifestFile, the order of the JSON lines in the AugmentedManifestFile is shuffled. The shuffling order is determined using the Seed value.  For Pipe input mode, when ShuffleConfig is specified shuffling is done at the start of every epoch. With large datasets, this ensures that the order of the training data is different for each epoch, and it helps reduce bias and possible overfitting. In a multi-node training job when ShuffleConfig is combined with S3DataDistributionType of ShardedByS3Key, the data is shuffled across nodes so that the content sent to a particular node on the first epoch might be sent to a different node on the second epoch. (see [below for nested schema](#nestedatt--spec--input_data_config--shuffle_config))

<a id="nestedatt--spec--input_data_config--data_source"></a>
### Nested Schema for `spec.input_data_config.data_source`

Optional:

- `file_system_data_source` (Attributes) Specifies a file system data source for a channel. (see [below for nested schema](#nestedatt--spec--input_data_config--data_source--file_system_data_source))
- `s3_data_source` (Attributes) Describes the S3 data source. (see [below for nested schema](#nestedatt--spec--input_data_config--data_source--s3_data_source))

<a id="nestedatt--spec--input_data_config--data_source--file_system_data_source"></a>
### Nested Schema for `spec.input_data_config.data_source.s3_data_source`

Optional:

- `directory_path` (String)
- `file_system_access_mode` (String)
- `file_system_id` (String)
- `file_system_type` (String)


<a id="nestedatt--spec--input_data_config--data_source--s3_data_source"></a>
### Nested Schema for `spec.input_data_config.data_source.s3_data_source`

Optional:

- `attribute_names` (List of String)
- `s3_data_distribution_type` (String)
- `s3_data_type` (String)
- `s3_uri` (String)



<a id="nestedatt--spec--input_data_config--shuffle_config"></a>
### Nested Schema for `spec.input_data_config.shuffle_config`

Optional:

- `seed` (Number)



<a id="nestedatt--spec--profiler_config"></a>
### Nested Schema for `spec.profiler_config`

Optional:

- `profiling_interval_in_milliseconds` (Number)
- `profiling_parameters` (Map of String)
- `s3_output_path` (String)


<a id="nestedatt--spec--profiler_rule_configurations"></a>
### Nested Schema for `spec.profiler_rule_configurations`

Optional:

- `instance_type` (String)
- `local_path` (String)
- `rule_configuration_name` (String)
- `rule_evaluator_image` (String)
- `rule_parameters` (Map of String)
- `s3_output_path` (String)
- `volume_size_in_gb` (Number)


<a id="nestedatt--spec--tags"></a>
### Nested Schema for `spec.tags`

Optional:

- `key` (String)
- `value` (String)


<a id="nestedatt--spec--tensor_board_output_config"></a>
### Nested Schema for `spec.tensor_board_output_config`

Optional:

- `local_path` (String)
- `s3_output_path` (String)


<a id="nestedatt--spec--vpc_config"></a>
### Nested Schema for `spec.vpc_config`

Optional:

- `security_group_i_ds` (List of String)
- `subnets` (List of String)

