/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"gopkg.in/yaml.v3"
	"time"
)

type EmrcontainersServicesK8SAwsJobRunV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*EmrcontainersServicesK8SAwsJobRunV1Alpha1Resource)(nil)
)

type EmrcontainersServicesK8SAwsJobRunV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type EmrcontainersServicesK8SAwsJobRunV1Alpha1GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		ConfigurationOverrides *string `tfsdk:"configuration_overrides" yaml:"configurationOverrides,omitempty"`

		ExecutionRoleARN *string `tfsdk:"execution_role_arn" yaml:"executionRoleARN,omitempty"`

		JobDriver *struct {
			SparkSubmitJobDriver *struct {
				EntryPoint *string `tfsdk:"entry_point" yaml:"entryPoint,omitempty"`

				EntryPointArguments *[]string `tfsdk:"entry_point_arguments" yaml:"entryPointArguments,omitempty"`

				SparkSubmitParameters *string `tfsdk:"spark_submit_parameters" yaml:"sparkSubmitParameters,omitempty"`
			} `tfsdk:"spark_submit_job_driver" yaml:"sparkSubmitJobDriver,omitempty"`
		} `tfsdk:"job_driver" yaml:"jobDriver,omitempty"`

		Name *string `tfsdk:"name" yaml:"name,omitempty"`

		ReleaseLabel *string `tfsdk:"release_label" yaml:"releaseLabel,omitempty"`

		Tags *map[string]string `tfsdk:"tags" yaml:"tags,omitempty"`

		VirtualClusterID *string `tfsdk:"virtual_cluster_id" yaml:"virtualClusterID,omitempty"`

		VirtualClusterRef *struct {
			From *struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`
			} `tfsdk:"from" yaml:"from,omitempty"`
		} `tfsdk:"virtual_cluster_ref" yaml:"virtualClusterRef,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewEmrcontainersServicesK8SAwsJobRunV1Alpha1Resource() resource.Resource {
	return &EmrcontainersServicesK8SAwsJobRunV1Alpha1Resource{}
}

func (r *EmrcontainersServicesK8SAwsJobRunV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_emrcontainers_services_k8s_aws_job_run_v1alpha1"
}

func (r *EmrcontainersServicesK8SAwsJobRunV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "JobRun is the Schema for the JobRuns API",
		MarkdownDescription: "JobRun is the Schema for the JobRuns API",
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Description:         "The timestamp of the last change to this resource.",
				MarkdownDescription: "The timestamp of the last change to this resource.",
				Type:                types.Int64Type,
				Computed:            true,
				Optional:            false,
			},

			"yaml": {
				Description:         "The generated manifest in YAML format.",
				MarkdownDescription: "The generated manifest in YAML format.",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"metadata": {
				Description:         "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				MarkdownDescription: "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				Required:            true,
				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
					"name": {
						Description:         "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						MarkdownDescription: "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						Type:                types.StringType,
						Required:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.NameValidator(),
						},
					},

					"namespace": {
						Description:         "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						MarkdownDescription: "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						Type:                types.StringType,
						Optional:            true,
					},

					"labels": {
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.LabelValidator(),
						},
					},
					"annotations": {
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.AnnotationValidator(),
						},
					},
				}),
			},

			"api_version": {
				Description:         "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				MarkdownDescription: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"kind": {
				Description:         "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				MarkdownDescription: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"spec": {
				Description:         "JobRunSpec defines the desired state of JobRun.  This entity describes a job run. A job run is a unit of work, such as a Spark jar, PySpark script, or SparkSQL query, that you submit to Amazon EMR on EKS.",
				MarkdownDescription: "JobRunSpec defines the desired state of JobRun.  This entity describes a job run. A job run is a unit of work, such as a Spark jar, PySpark script, or SparkSQL query, that you submit to Amazon EMR on EKS.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"configuration_overrides": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"execution_role_arn": {
						Description:         "The execution role ARN for the job run.",
						MarkdownDescription: "The execution role ARN for the job run.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"job_driver": {
						Description:         "The job driver for the job run.",
						MarkdownDescription: "The job driver for the job run.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"spark_submit_job_driver": {
								Description:         "The information about job driver for Spark submit.",
								MarkdownDescription: "The information about job driver for Spark submit.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"entry_point": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"entry_point_arguments": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"spark_submit_parameters": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: true,
						Optional: false,
						Computed: false,
					},

					"name": {
						Description:         "The name of the job run.",
						MarkdownDescription: "The name of the job run.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"release_label": {
						Description:         "The Amazon EMR release version to use for the job run.",
						MarkdownDescription: "The Amazon EMR release version to use for the job run.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"tags": {
						Description:         "The tags assigned to job runs.",
						MarkdownDescription: "The tags assigned to job runs.",

						Type: types.MapType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"virtual_cluster_id": {
						Description:         "The virtual cluster ID for which the job run request is submitted.",
						MarkdownDescription: "The virtual cluster ID for which the job run request is submitted.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"virtual_cluster_ref": {
						Description:         "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReference type to provide more user friendly syntax for references using 'from' field Ex: APIIDRef: from: name: my-api",
						MarkdownDescription: "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReference type to provide more user friendly syntax for references using 'from' field Ex: APIIDRef: from: name: my-api",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"from": {
								Description:         "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",
								MarkdownDescription: "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},
				}),

				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}, nil
}

func (r *EmrcontainersServicesK8SAwsJobRunV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_emrcontainers_services_k8s_aws_job_run_v1alpha1")

	var state EmrcontainersServicesK8SAwsJobRunV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel EmrcontainersServicesK8SAwsJobRunV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("emrcontainers.services.k8s.aws/v1alpha1")
	goModel.Kind = utilities.Ptr("JobRun")

	state.Id = types.Int64{Value: time.Now().UnixNano()}
	state.ApiVersion = types.String{Value: *goModel.ApiVersion}
	state.Kind = types.String{Value: *goModel.Kind}

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.String{Value: string(marshal)}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *EmrcontainersServicesK8SAwsJobRunV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_emrcontainers_services_k8s_aws_job_run_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *EmrcontainersServicesK8SAwsJobRunV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_emrcontainers_services_k8s_aws_job_run_v1alpha1")

	var state EmrcontainersServicesK8SAwsJobRunV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel EmrcontainersServicesK8SAwsJobRunV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("emrcontainers.services.k8s.aws/v1alpha1")
	goModel.Kind = utilities.Ptr("JobRun")

	state.Id = types.Int64{Value: time.Now().UnixNano()}
	state.ApiVersion = types.String{Value: *goModel.ApiVersion}
	state.Kind = types.String{Value: *goModel.Kind}

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.String{Value: string(marshal)}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *EmrcontainersServicesK8SAwsJobRunV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_emrcontainers_services_k8s_aws_job_run_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
