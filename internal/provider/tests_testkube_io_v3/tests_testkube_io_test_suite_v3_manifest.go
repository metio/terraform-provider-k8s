/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package tests_testkube_io_v3

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
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &TestsTestkubeIoTestSuiteV3Manifest{}
)

func NewTestsTestkubeIoTestSuiteV3Manifest() datasource.DataSource {
	return &TestsTestkubeIoTestSuiteV3Manifest{}
}

type TestsTestkubeIoTestSuiteV3Manifest struct{}

type TestsTestkubeIoTestSuiteV3ManifestData struct {
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
		After *[]struct {
			DownloadArtifacts *struct {
				AllPreviousSteps    *bool     `tfsdk:"all_previous_steps" json:"allPreviousSteps,omitempty"`
				PreviousStepNumbers *[]string `tfsdk:"previous_step_numbers" json:"previousStepNumbers,omitempty"`
				PreviousTestNames   *[]string `tfsdk:"previous_test_names" json:"previousTestNames,omitempty"`
			} `tfsdk:"download_artifacts" json:"downloadArtifacts,omitempty"`
			Execute *[]struct {
				Delay            *string `tfsdk:"delay" json:"delay,omitempty"`
				ExecutionRequest *struct {
					Args                     *[]string          `tfsdk:"args" json:"args,omitempty"`
					ArgsMode                 *string            `tfsdk:"args_mode" json:"argsMode,omitempty"`
					Command                  *[]string          `tfsdk:"command" json:"command,omitempty"`
					CronJobTemplate          *string            `tfsdk:"cron_job_template" json:"cronJobTemplate,omitempty"`
					CronJobTemplateReference *string            `tfsdk:"cron_job_template_reference" json:"cronJobTemplateReference,omitempty"`
					DisableWebhooks          *bool              `tfsdk:"disable_webhooks" json:"disableWebhooks,omitempty"`
					ExecutionLabels          *map[string]string `tfsdk:"execution_labels" json:"executionLabels,omitempty"`
					HttpProxy                *string            `tfsdk:"http_proxy" json:"httpProxy,omitempty"`
					HttpsProxy               *string            `tfsdk:"https_proxy" json:"httpsProxy,omitempty"`
					JobTemplate              *string            `tfsdk:"job_template" json:"jobTemplate,omitempty"`
					JobTemplateReference     *string            `tfsdk:"job_template_reference" json:"jobTemplateReference,omitempty"`
					NegativeTest             *bool              `tfsdk:"negative_test" json:"negativeTest,omitempty"`
					PvcTemplate              *string            `tfsdk:"pvc_template" json:"pvcTemplate,omitempty"`
					PvcTemplateReference     *string            `tfsdk:"pvc_template_reference" json:"pvcTemplateReference,omitempty"`
					RunningContext           *struct {
						Context *string `tfsdk:"context" json:"context,omitempty"`
						Type    *string `tfsdk:"type" json:"type,omitempty"`
					} `tfsdk:"running_context" json:"runningContext,omitempty"`
					ScraperTemplate          *string `tfsdk:"scraper_template" json:"scraperTemplate,omitempty"`
					ScraperTemplateReference *string `tfsdk:"scraper_template_reference" json:"scraperTemplateReference,omitempty"`
					Sync                     *bool   `tfsdk:"sync" json:"sync,omitempty"`
					Variables                *struct {
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Type      *string `tfsdk:"type" json:"type,omitempty"`
						Value     *string `tfsdk:"value" json:"value,omitempty"`
						ValueFrom *struct {
							ConfigMapKeyRef *struct {
								Key      *string `tfsdk:"key" json:"key,omitempty"`
								Name     *string `tfsdk:"name" json:"name,omitempty"`
								Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
							} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
							FieldRef *struct {
								ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
								FieldPath  *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
							} `tfsdk:"field_ref" json:"fieldRef,omitempty"`
							ResourceFieldRef *struct {
								ContainerName *string `tfsdk:"container_name" json:"containerName,omitempty"`
								Divisor       *string `tfsdk:"divisor" json:"divisor,omitempty"`
								Resource      *string `tfsdk:"resource" json:"resource,omitempty"`
							} `tfsdk:"resource_field_ref" json:"resourceFieldRef,omitempty"`
							SecretKeyRef *struct {
								Key      *string `tfsdk:"key" json:"key,omitempty"`
								Name     *string `tfsdk:"name" json:"name,omitempty"`
								Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
							} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
						} `tfsdk:"value_from" json:"valueFrom,omitempty"`
					} `tfsdk:"variables" json:"variables,omitempty"`
				} `tfsdk:"execution_request" json:"executionRequest,omitempty"`
				Test *string `tfsdk:"test" json:"test,omitempty"`
			} `tfsdk:"execute" json:"execute,omitempty"`
			StopOnFailure *bool `tfsdk:"stop_on_failure" json:"stopOnFailure,omitempty"`
		} `tfsdk:"after" json:"after,omitempty"`
		Before *[]struct {
			DownloadArtifacts *struct {
				AllPreviousSteps    *bool     `tfsdk:"all_previous_steps" json:"allPreviousSteps,omitempty"`
				PreviousStepNumbers *[]string `tfsdk:"previous_step_numbers" json:"previousStepNumbers,omitempty"`
				PreviousTestNames   *[]string `tfsdk:"previous_test_names" json:"previousTestNames,omitempty"`
			} `tfsdk:"download_artifacts" json:"downloadArtifacts,omitempty"`
			Execute *[]struct {
				Delay            *string `tfsdk:"delay" json:"delay,omitempty"`
				ExecutionRequest *struct {
					Args                     *[]string          `tfsdk:"args" json:"args,omitempty"`
					ArgsMode                 *string            `tfsdk:"args_mode" json:"argsMode,omitempty"`
					Command                  *[]string          `tfsdk:"command" json:"command,omitempty"`
					CronJobTemplate          *string            `tfsdk:"cron_job_template" json:"cronJobTemplate,omitempty"`
					CronJobTemplateReference *string            `tfsdk:"cron_job_template_reference" json:"cronJobTemplateReference,omitempty"`
					DisableWebhooks          *bool              `tfsdk:"disable_webhooks" json:"disableWebhooks,omitempty"`
					ExecutionLabels          *map[string]string `tfsdk:"execution_labels" json:"executionLabels,omitempty"`
					HttpProxy                *string            `tfsdk:"http_proxy" json:"httpProxy,omitempty"`
					HttpsProxy               *string            `tfsdk:"https_proxy" json:"httpsProxy,omitempty"`
					JobTemplate              *string            `tfsdk:"job_template" json:"jobTemplate,omitempty"`
					JobTemplateReference     *string            `tfsdk:"job_template_reference" json:"jobTemplateReference,omitempty"`
					NegativeTest             *bool              `tfsdk:"negative_test" json:"negativeTest,omitempty"`
					PvcTemplate              *string            `tfsdk:"pvc_template" json:"pvcTemplate,omitempty"`
					PvcTemplateReference     *string            `tfsdk:"pvc_template_reference" json:"pvcTemplateReference,omitempty"`
					RunningContext           *struct {
						Context *string `tfsdk:"context" json:"context,omitempty"`
						Type    *string `tfsdk:"type" json:"type,omitempty"`
					} `tfsdk:"running_context" json:"runningContext,omitempty"`
					ScraperTemplate          *string `tfsdk:"scraper_template" json:"scraperTemplate,omitempty"`
					ScraperTemplateReference *string `tfsdk:"scraper_template_reference" json:"scraperTemplateReference,omitempty"`
					Sync                     *bool   `tfsdk:"sync" json:"sync,omitempty"`
					Variables                *struct {
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Type      *string `tfsdk:"type" json:"type,omitempty"`
						Value     *string `tfsdk:"value" json:"value,omitempty"`
						ValueFrom *struct {
							ConfigMapKeyRef *struct {
								Key      *string `tfsdk:"key" json:"key,omitempty"`
								Name     *string `tfsdk:"name" json:"name,omitempty"`
								Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
							} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
							FieldRef *struct {
								ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
								FieldPath  *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
							} `tfsdk:"field_ref" json:"fieldRef,omitempty"`
							ResourceFieldRef *struct {
								ContainerName *string `tfsdk:"container_name" json:"containerName,omitempty"`
								Divisor       *string `tfsdk:"divisor" json:"divisor,omitempty"`
								Resource      *string `tfsdk:"resource" json:"resource,omitempty"`
							} `tfsdk:"resource_field_ref" json:"resourceFieldRef,omitempty"`
							SecretKeyRef *struct {
								Key      *string `tfsdk:"key" json:"key,omitempty"`
								Name     *string `tfsdk:"name" json:"name,omitempty"`
								Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
							} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
						} `tfsdk:"value_from" json:"valueFrom,omitempty"`
					} `tfsdk:"variables" json:"variables,omitempty"`
				} `tfsdk:"execution_request" json:"executionRequest,omitempty"`
				Test *string `tfsdk:"test" json:"test,omitempty"`
			} `tfsdk:"execute" json:"execute,omitempty"`
			StopOnFailure *bool `tfsdk:"stop_on_failure" json:"stopOnFailure,omitempty"`
		} `tfsdk:"before" json:"before,omitempty"`
		Description      *string `tfsdk:"description" json:"description,omitempty"`
		ExecutionRequest *struct {
			CronJobTemplate          *string            `tfsdk:"cron_job_template" json:"cronJobTemplate,omitempty"`
			CronJobTemplateReference *string            `tfsdk:"cron_job_template_reference" json:"cronJobTemplateReference,omitempty"`
			DisableWebhooks          *bool              `tfsdk:"disable_webhooks" json:"disableWebhooks,omitempty"`
			ExecutionLabels          *map[string]string `tfsdk:"execution_labels" json:"executionLabels,omitempty"`
			HttpProxy                *string            `tfsdk:"http_proxy" json:"httpProxy,omitempty"`
			HttpsProxy               *string            `tfsdk:"https_proxy" json:"httpsProxy,omitempty"`
			JobTemplate              *string            `tfsdk:"job_template" json:"jobTemplate,omitempty"`
			JobTemplateReference     *string            `tfsdk:"job_template_reference" json:"jobTemplateReference,omitempty"`
			Labels                   *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			Name                     *string            `tfsdk:"name" json:"name,omitempty"`
			Namespace                *string            `tfsdk:"namespace" json:"namespace,omitempty"`
			PvcTemplate              *string            `tfsdk:"pvc_template" json:"pvcTemplate,omitempty"`
			PvcTemplateReference     *string            `tfsdk:"pvc_template_reference" json:"pvcTemplateReference,omitempty"`
			ScraperTemplate          *string            `tfsdk:"scraper_template" json:"scraperTemplate,omitempty"`
			ScraperTemplateReference *string            `tfsdk:"scraper_template_reference" json:"scraperTemplateReference,omitempty"`
			SecretUUID               *string            `tfsdk:"secret_uuid" json:"secretUUID,omitempty"`
			Sync                     *bool              `tfsdk:"sync" json:"sync,omitempty"`
			Timeout                  *int64             `tfsdk:"timeout" json:"timeout,omitempty"`
			Variables                *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Type      *string `tfsdk:"type" json:"type,omitempty"`
				Value     *string `tfsdk:"value" json:"value,omitempty"`
				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
					} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
					FieldRef *struct {
						ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
						FieldPath  *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
					} `tfsdk:"field_ref" json:"fieldRef,omitempty"`
					ResourceFieldRef *struct {
						ContainerName *string `tfsdk:"container_name" json:"containerName,omitempty"`
						Divisor       *string `tfsdk:"divisor" json:"divisor,omitempty"`
						Resource      *string `tfsdk:"resource" json:"resource,omitempty"`
					} `tfsdk:"resource_field_ref" json:"resourceFieldRef,omitempty"`
					SecretKeyRef *struct {
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"variables" json:"variables,omitempty"`
		} `tfsdk:"execution_request" json:"executionRequest,omitempty"`
		Repeats  *int64  `tfsdk:"repeats" json:"repeats,omitempty"`
		Schedule *string `tfsdk:"schedule" json:"schedule,omitempty"`
		Steps    *[]struct {
			DownloadArtifacts *struct {
				AllPreviousSteps    *bool     `tfsdk:"all_previous_steps" json:"allPreviousSteps,omitempty"`
				PreviousStepNumbers *[]string `tfsdk:"previous_step_numbers" json:"previousStepNumbers,omitempty"`
				PreviousTestNames   *[]string `tfsdk:"previous_test_names" json:"previousTestNames,omitempty"`
			} `tfsdk:"download_artifacts" json:"downloadArtifacts,omitempty"`
			Execute *[]struct {
				Delay            *string `tfsdk:"delay" json:"delay,omitempty"`
				ExecutionRequest *struct {
					Args                     *[]string          `tfsdk:"args" json:"args,omitempty"`
					ArgsMode                 *string            `tfsdk:"args_mode" json:"argsMode,omitempty"`
					Command                  *[]string          `tfsdk:"command" json:"command,omitempty"`
					CronJobTemplate          *string            `tfsdk:"cron_job_template" json:"cronJobTemplate,omitempty"`
					CronJobTemplateReference *string            `tfsdk:"cron_job_template_reference" json:"cronJobTemplateReference,omitempty"`
					DisableWebhooks          *bool              `tfsdk:"disable_webhooks" json:"disableWebhooks,omitempty"`
					ExecutionLabels          *map[string]string `tfsdk:"execution_labels" json:"executionLabels,omitempty"`
					HttpProxy                *string            `tfsdk:"http_proxy" json:"httpProxy,omitempty"`
					HttpsProxy               *string            `tfsdk:"https_proxy" json:"httpsProxy,omitempty"`
					JobTemplate              *string            `tfsdk:"job_template" json:"jobTemplate,omitempty"`
					JobTemplateReference     *string            `tfsdk:"job_template_reference" json:"jobTemplateReference,omitempty"`
					NegativeTest             *bool              `tfsdk:"negative_test" json:"negativeTest,omitempty"`
					PvcTemplate              *string            `tfsdk:"pvc_template" json:"pvcTemplate,omitempty"`
					PvcTemplateReference     *string            `tfsdk:"pvc_template_reference" json:"pvcTemplateReference,omitempty"`
					RunningContext           *struct {
						Context *string `tfsdk:"context" json:"context,omitempty"`
						Type    *string `tfsdk:"type" json:"type,omitempty"`
					} `tfsdk:"running_context" json:"runningContext,omitempty"`
					ScraperTemplate          *string `tfsdk:"scraper_template" json:"scraperTemplate,omitempty"`
					ScraperTemplateReference *string `tfsdk:"scraper_template_reference" json:"scraperTemplateReference,omitempty"`
					Sync                     *bool   `tfsdk:"sync" json:"sync,omitempty"`
					Variables                *struct {
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Type      *string `tfsdk:"type" json:"type,omitempty"`
						Value     *string `tfsdk:"value" json:"value,omitempty"`
						ValueFrom *struct {
							ConfigMapKeyRef *struct {
								Key      *string `tfsdk:"key" json:"key,omitempty"`
								Name     *string `tfsdk:"name" json:"name,omitempty"`
								Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
							} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
							FieldRef *struct {
								ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
								FieldPath  *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
							} `tfsdk:"field_ref" json:"fieldRef,omitempty"`
							ResourceFieldRef *struct {
								ContainerName *string `tfsdk:"container_name" json:"containerName,omitempty"`
								Divisor       *string `tfsdk:"divisor" json:"divisor,omitempty"`
								Resource      *string `tfsdk:"resource" json:"resource,omitempty"`
							} `tfsdk:"resource_field_ref" json:"resourceFieldRef,omitempty"`
							SecretKeyRef *struct {
								Key      *string `tfsdk:"key" json:"key,omitempty"`
								Name     *string `tfsdk:"name" json:"name,omitempty"`
								Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
							} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
						} `tfsdk:"value_from" json:"valueFrom,omitempty"`
					} `tfsdk:"variables" json:"variables,omitempty"`
				} `tfsdk:"execution_request" json:"executionRequest,omitempty"`
				Test *string `tfsdk:"test" json:"test,omitempty"`
			} `tfsdk:"execute" json:"execute,omitempty"`
			StopOnFailure *bool `tfsdk:"stop_on_failure" json:"stopOnFailure,omitempty"`
		} `tfsdk:"steps" json:"steps,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *TestsTestkubeIoTestSuiteV3Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_tests_testkube_io_test_suite_v3_manifest"
}

func (r *TestsTestkubeIoTestSuiteV3Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "TestSuite is the Schema for the testsuites API",
		MarkdownDescription: "TestSuite is the Schema for the testsuites API",
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
				Description:         "TestSuiteSpec defines the desired state of TestSuite",
				MarkdownDescription: "TestSuiteSpec defines the desired state of TestSuite",
				Attributes: map[string]schema.Attribute{
					"after": schema.ListNestedAttribute{
						Description:         "After batch steps is list of batch tests which will be sequentially orchestrated for parallel tests in each batch",
						MarkdownDescription: "After batch steps is list of batch tests which will be sequentially orchestrated for parallel tests in each batch",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"download_artifacts": schema.SingleNestedAttribute{
									Description:         "options to download artifacts from previous steps",
									MarkdownDescription: "options to download artifacts from previous steps",
									Attributes: map[string]schema.Attribute{
										"all_previous_steps": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"previous_step_numbers": schema.ListAttribute{
											Description:         "previous step numbers starting from 1",
											MarkdownDescription: "previous step numbers starting from 1",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"previous_test_names": schema.ListAttribute{
											Description:         "previous test names",
											MarkdownDescription: "previous test names",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"execute": schema.ListNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"delay": schema.StringAttribute{
												Description:         "delay duration in time units",
												MarkdownDescription: "delay duration in time units",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"execution_request": schema.SingleNestedAttribute{
												Description:         "TestSuiteStepExecutionRequest contains parameters to be used by the executions. These fields will be passed to the execution when a Test Suite is queued for execution. TestSuiteStepExecutionRequest parameters have the highest priority. They override the values coming from Test Suites, Tests, and Test Executions.",
												MarkdownDescription: "TestSuiteStepExecutionRequest contains parameters to be used by the executions. These fields will be passed to the execution when a Test Suite is queued for execution. TestSuiteStepExecutionRequest parameters have the highest priority. They override the values coming from Test Suites, Tests, and Test Executions.",
												Attributes: map[string]schema.Attribute{
													"args": schema.ListAttribute{
														Description:         "additional executor binary arguments",
														MarkdownDescription: "additional executor binary arguments",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"args_mode": schema.StringAttribute{
														Description:         "usage mode for arguments",
														MarkdownDescription: "usage mode for arguments",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("append", "override", "replace"),
														},
													},

													"command": schema.ListAttribute{
														Description:         "executor binary command",
														MarkdownDescription: "executor binary command",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"cron_job_template": schema.StringAttribute{
														Description:         "cron job template extensions",
														MarkdownDescription: "cron job template extensions",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"cron_job_template_reference": schema.StringAttribute{
														Description:         "cron job template extensions reference",
														MarkdownDescription: "cron job template extensions reference",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"disable_webhooks": schema.BoolAttribute{
														Description:         "whether webhooks should be called on execution Deprecated: field is not used",
														MarkdownDescription: "whether webhooks should be called on execution Deprecated: field is not used",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"execution_labels": schema.MapAttribute{
														Description:         "test execution labels",
														MarkdownDescription: "test execution labels",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"http_proxy": schema.StringAttribute{
														Description:         "http proxy for executor containers",
														MarkdownDescription: "http proxy for executor containers",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"https_proxy": schema.StringAttribute{
														Description:         "https proxy for executor containers",
														MarkdownDescription: "https proxy for executor containers",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"job_template": schema.StringAttribute{
														Description:         "job template extensions",
														MarkdownDescription: "job template extensions",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"job_template_reference": schema.StringAttribute{
														Description:         "job template extensions reference",
														MarkdownDescription: "job template extensions reference",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"negative_test": schema.BoolAttribute{
														Description:         "negative test will fail the execution if it is a success and it will succeed if it is a failure",
														MarkdownDescription: "negative test will fail the execution if it is a success and it will succeed if it is a failure",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"pvc_template": schema.StringAttribute{
														Description:         "pvc template extensions",
														MarkdownDescription: "pvc template extensions",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"pvc_template_reference": schema.StringAttribute{
														Description:         "pvc template extensions reference",
														MarkdownDescription: "pvc template extensions reference",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"running_context": schema.SingleNestedAttribute{
														Description:         "RunningContext for test or test suite execution",
														MarkdownDescription: "RunningContext for test or test suite execution",
														Attributes: map[string]schema.Attribute{
															"context": schema.StringAttribute{
																Description:         "Context value depending from its type",
																MarkdownDescription: "Context value depending from its type",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"type": schema.StringAttribute{
																Description:         "One of possible context types",
																MarkdownDescription: "One of possible context types",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"scraper_template": schema.StringAttribute{
														Description:         "scraper template extensions",
														MarkdownDescription: "scraper template extensions",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"scraper_template_reference": schema.StringAttribute{
														Description:         "scraper template extensions reference",
														MarkdownDescription: "scraper template extensions reference",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"sync": schema.BoolAttribute{
														Description:         "whether to start execution sync or async",
														MarkdownDescription: "whether to start execution sync or async",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"variables": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "variable name",
																MarkdownDescription: "variable name",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"type": schema.StringAttribute{
																Description:         "variable type",
																MarkdownDescription: "variable type",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"value": schema.StringAttribute{
																Description:         "variable string value",
																MarkdownDescription: "variable string value",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"value_from": schema.SingleNestedAttribute{
																Description:         "or load it from var source",
																MarkdownDescription: "or load it from var source",
																Attributes: map[string]schema.Attribute{
																	"config_map_key_ref": schema.SingleNestedAttribute{
																		Description:         "Selects a key of a ConfigMap.",
																		MarkdownDescription: "Selects a key of a ConfigMap.",
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "The key to select.",
																				MarkdownDescription: "The key to select.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"optional": schema.BoolAttribute{
																				Description:         "Specify whether the ConfigMap or its key must be defined",
																				MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},
																		},
																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"field_ref": schema.SingleNestedAttribute{
																		Description:         "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
																		MarkdownDescription: "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
																		Attributes: map[string]schema.Attribute{
																			"api_version": schema.StringAttribute{
																				Description:         "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
																				MarkdownDescription: "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"field_path": schema.StringAttribute{
																				Description:         "Path of the field to select in the specified API version.",
																				MarkdownDescription: "Path of the field to select in the specified API version.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},
																		},
																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"resource_field_ref": schema.SingleNestedAttribute{
																		Description:         "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
																		MarkdownDescription: "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
																		Attributes: map[string]schema.Attribute{
																			"container_name": schema.StringAttribute{
																				Description:         "Container name: required for volumes, optional for env vars",
																				MarkdownDescription: "Container name: required for volumes, optional for env vars",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"divisor": schema.StringAttribute{
																				Description:         "Specifies the output format of the exposed resources, defaults to '1'",
																				MarkdownDescription: "Specifies the output format of the exposed resources, defaults to '1'",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"resource": schema.StringAttribute{
																				Description:         "Required: resource to select",
																				MarkdownDescription: "Required: resource to select",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},
																		},
																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"secret_key_ref": schema.SingleNestedAttribute{
																		Description:         "Selects a key of a secret in the pod's namespace",
																		MarkdownDescription: "Selects a key of a secret in the pod's namespace",
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "The key of the secret to select from. Must be a valid secret key.",
																				MarkdownDescription: "The key of the secret to select from. Must be a valid secret key.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"optional": schema.BoolAttribute{
																				Description:         "Specify whether the Secret or its key must be defined",
																				MarkdownDescription: "Specify whether the Secret or its key must be defined",
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
														Required: false,
														Optional: true,
														Computed: false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"test": schema.StringAttribute{
												Description:         "object name",
												MarkdownDescription: "object name",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"stop_on_failure": schema.BoolAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"before": schema.ListNestedAttribute{
						Description:         "Before batch steps is list of batch tests which will be sequentially orchestrated for parallel tests in each batch",
						MarkdownDescription: "Before batch steps is list of batch tests which will be sequentially orchestrated for parallel tests in each batch",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"download_artifacts": schema.SingleNestedAttribute{
									Description:         "options to download artifacts from previous steps",
									MarkdownDescription: "options to download artifacts from previous steps",
									Attributes: map[string]schema.Attribute{
										"all_previous_steps": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"previous_step_numbers": schema.ListAttribute{
											Description:         "previous step numbers starting from 1",
											MarkdownDescription: "previous step numbers starting from 1",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"previous_test_names": schema.ListAttribute{
											Description:         "previous test names",
											MarkdownDescription: "previous test names",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"execute": schema.ListNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"delay": schema.StringAttribute{
												Description:         "delay duration in time units",
												MarkdownDescription: "delay duration in time units",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"execution_request": schema.SingleNestedAttribute{
												Description:         "TestSuiteStepExecutionRequest contains parameters to be used by the executions. These fields will be passed to the execution when a Test Suite is queued for execution. TestSuiteStepExecutionRequest parameters have the highest priority. They override the values coming from Test Suites, Tests, and Test Executions.",
												MarkdownDescription: "TestSuiteStepExecutionRequest contains parameters to be used by the executions. These fields will be passed to the execution when a Test Suite is queued for execution. TestSuiteStepExecutionRequest parameters have the highest priority. They override the values coming from Test Suites, Tests, and Test Executions.",
												Attributes: map[string]schema.Attribute{
													"args": schema.ListAttribute{
														Description:         "additional executor binary arguments",
														MarkdownDescription: "additional executor binary arguments",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"args_mode": schema.StringAttribute{
														Description:         "usage mode for arguments",
														MarkdownDescription: "usage mode for arguments",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("append", "override", "replace"),
														},
													},

													"command": schema.ListAttribute{
														Description:         "executor binary command",
														MarkdownDescription: "executor binary command",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"cron_job_template": schema.StringAttribute{
														Description:         "cron job template extensions",
														MarkdownDescription: "cron job template extensions",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"cron_job_template_reference": schema.StringAttribute{
														Description:         "cron job template extensions reference",
														MarkdownDescription: "cron job template extensions reference",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"disable_webhooks": schema.BoolAttribute{
														Description:         "whether webhooks should be called on execution Deprecated: field is not used",
														MarkdownDescription: "whether webhooks should be called on execution Deprecated: field is not used",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"execution_labels": schema.MapAttribute{
														Description:         "test execution labels",
														MarkdownDescription: "test execution labels",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"http_proxy": schema.StringAttribute{
														Description:         "http proxy for executor containers",
														MarkdownDescription: "http proxy for executor containers",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"https_proxy": schema.StringAttribute{
														Description:         "https proxy for executor containers",
														MarkdownDescription: "https proxy for executor containers",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"job_template": schema.StringAttribute{
														Description:         "job template extensions",
														MarkdownDescription: "job template extensions",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"job_template_reference": schema.StringAttribute{
														Description:         "job template extensions reference",
														MarkdownDescription: "job template extensions reference",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"negative_test": schema.BoolAttribute{
														Description:         "negative test will fail the execution if it is a success and it will succeed if it is a failure",
														MarkdownDescription: "negative test will fail the execution if it is a success and it will succeed if it is a failure",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"pvc_template": schema.StringAttribute{
														Description:         "pvc template extensions",
														MarkdownDescription: "pvc template extensions",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"pvc_template_reference": schema.StringAttribute{
														Description:         "pvc template extensions reference",
														MarkdownDescription: "pvc template extensions reference",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"running_context": schema.SingleNestedAttribute{
														Description:         "RunningContext for test or test suite execution",
														MarkdownDescription: "RunningContext for test or test suite execution",
														Attributes: map[string]schema.Attribute{
															"context": schema.StringAttribute{
																Description:         "Context value depending from its type",
																MarkdownDescription: "Context value depending from its type",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"type": schema.StringAttribute{
																Description:         "One of possible context types",
																MarkdownDescription: "One of possible context types",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"scraper_template": schema.StringAttribute{
														Description:         "scraper template extensions",
														MarkdownDescription: "scraper template extensions",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"scraper_template_reference": schema.StringAttribute{
														Description:         "scraper template extensions reference",
														MarkdownDescription: "scraper template extensions reference",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"sync": schema.BoolAttribute{
														Description:         "whether to start execution sync or async",
														MarkdownDescription: "whether to start execution sync or async",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"variables": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "variable name",
																MarkdownDescription: "variable name",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"type": schema.StringAttribute{
																Description:         "variable type",
																MarkdownDescription: "variable type",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"value": schema.StringAttribute{
																Description:         "variable string value",
																MarkdownDescription: "variable string value",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"value_from": schema.SingleNestedAttribute{
																Description:         "or load it from var source",
																MarkdownDescription: "or load it from var source",
																Attributes: map[string]schema.Attribute{
																	"config_map_key_ref": schema.SingleNestedAttribute{
																		Description:         "Selects a key of a ConfigMap.",
																		MarkdownDescription: "Selects a key of a ConfigMap.",
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "The key to select.",
																				MarkdownDescription: "The key to select.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"optional": schema.BoolAttribute{
																				Description:         "Specify whether the ConfigMap or its key must be defined",
																				MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},
																		},
																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"field_ref": schema.SingleNestedAttribute{
																		Description:         "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
																		MarkdownDescription: "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
																		Attributes: map[string]schema.Attribute{
																			"api_version": schema.StringAttribute{
																				Description:         "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
																				MarkdownDescription: "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"field_path": schema.StringAttribute{
																				Description:         "Path of the field to select in the specified API version.",
																				MarkdownDescription: "Path of the field to select in the specified API version.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},
																		},
																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"resource_field_ref": schema.SingleNestedAttribute{
																		Description:         "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
																		MarkdownDescription: "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
																		Attributes: map[string]schema.Attribute{
																			"container_name": schema.StringAttribute{
																				Description:         "Container name: required for volumes, optional for env vars",
																				MarkdownDescription: "Container name: required for volumes, optional for env vars",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"divisor": schema.StringAttribute{
																				Description:         "Specifies the output format of the exposed resources, defaults to '1'",
																				MarkdownDescription: "Specifies the output format of the exposed resources, defaults to '1'",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"resource": schema.StringAttribute{
																				Description:         "Required: resource to select",
																				MarkdownDescription: "Required: resource to select",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},
																		},
																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"secret_key_ref": schema.SingleNestedAttribute{
																		Description:         "Selects a key of a secret in the pod's namespace",
																		MarkdownDescription: "Selects a key of a secret in the pod's namespace",
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "The key of the secret to select from. Must be a valid secret key.",
																				MarkdownDescription: "The key of the secret to select from. Must be a valid secret key.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"optional": schema.BoolAttribute{
																				Description:         "Specify whether the Secret or its key must be defined",
																				MarkdownDescription: "Specify whether the Secret or its key must be defined",
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
														Required: false,
														Optional: true,
														Computed: false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"test": schema.StringAttribute{
												Description:         "object name",
												MarkdownDescription: "object name",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"stop_on_failure": schema.BoolAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"description": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"execution_request": schema.SingleNestedAttribute{
						Description:         "test suite execution request body",
						MarkdownDescription: "test suite execution request body",
						Attributes: map[string]schema.Attribute{
							"cron_job_template": schema.StringAttribute{
								Description:         "cron job template extensions",
								MarkdownDescription: "cron job template extensions",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"cron_job_template_reference": schema.StringAttribute{
								Description:         "name of the template resource",
								MarkdownDescription: "name of the template resource",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"disable_webhooks": schema.BoolAttribute{
								Description:         "whether webhooks should be called on execution Deprecated: field is not used",
								MarkdownDescription: "whether webhooks should be called on execution Deprecated: field is not used",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"execution_labels": schema.MapAttribute{
								Description:         "execution labels",
								MarkdownDescription: "execution labels",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"http_proxy": schema.StringAttribute{
								Description:         "http proxy for executor containers",
								MarkdownDescription: "http proxy for executor containers",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"https_proxy": schema.StringAttribute{
								Description:         "https proxy for executor containers",
								MarkdownDescription: "https proxy for executor containers",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"job_template": schema.StringAttribute{
								Description:         "job template extensions",
								MarkdownDescription: "job template extensions",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"job_template_reference": schema.StringAttribute{
								Description:         "name of the template resource",
								MarkdownDescription: "name of the template resource",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"labels": schema.MapAttribute{
								Description:         "test suite labels",
								MarkdownDescription: "test suite labels",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "test execution custom name",
								MarkdownDescription: "test execution custom name",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"namespace": schema.StringAttribute{
								Description:         "test kubernetes namespace ('testkube' when not set)",
								MarkdownDescription: "test kubernetes namespace ('testkube' when not set)",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"pvc_template": schema.StringAttribute{
								Description:         "pvc template extensions",
								MarkdownDescription: "pvc template extensions",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"pvc_template_reference": schema.StringAttribute{
								Description:         "name of the template resource",
								MarkdownDescription: "name of the template resource",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"scraper_template": schema.StringAttribute{
								Description:         "scraper template extensions",
								MarkdownDescription: "scraper template extensions",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"scraper_template_reference": schema.StringAttribute{
								Description:         "name of the template resource",
								MarkdownDescription: "name of the template resource",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"secret_uuid": schema.StringAttribute{
								Description:         "secret uuid",
								MarkdownDescription: "secret uuid",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"sync": schema.BoolAttribute{
								Description:         "whether to start execution sync or async",
								MarkdownDescription: "whether to start execution sync or async",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"timeout": schema.Int64Attribute{
								Description:         "timeout for test suite execution",
								MarkdownDescription: "timeout for test suite execution",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"variables": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "variable name",
										MarkdownDescription: "variable name",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"type": schema.StringAttribute{
										Description:         "variable type",
										MarkdownDescription: "variable type",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"value": schema.StringAttribute{
										Description:         "variable string value",
										MarkdownDescription: "variable string value",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"value_from": schema.SingleNestedAttribute{
										Description:         "or load it from var source",
										MarkdownDescription: "or load it from var source",
										Attributes: map[string]schema.Attribute{
											"config_map_key_ref": schema.SingleNestedAttribute{
												Description:         "Selects a key of a ConfigMap.",
												MarkdownDescription: "Selects a key of a ConfigMap.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key to select.",
														MarkdownDescription: "The key to select.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
														MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"optional": schema.BoolAttribute{
														Description:         "Specify whether the ConfigMap or its key must be defined",
														MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"field_ref": schema.SingleNestedAttribute{
												Description:         "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
												MarkdownDescription: "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
												Attributes: map[string]schema.Attribute{
													"api_version": schema.StringAttribute{
														Description:         "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
														MarkdownDescription: "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"field_path": schema.StringAttribute{
														Description:         "Path of the field to select in the specified API version.",
														MarkdownDescription: "Path of the field to select in the specified API version.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"resource_field_ref": schema.SingleNestedAttribute{
												Description:         "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
												MarkdownDescription: "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
												Attributes: map[string]schema.Attribute{
													"container_name": schema.StringAttribute{
														Description:         "Container name: required for volumes, optional for env vars",
														MarkdownDescription: "Container name: required for volumes, optional for env vars",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"divisor": schema.StringAttribute{
														Description:         "Specifies the output format of the exposed resources, defaults to '1'",
														MarkdownDescription: "Specifies the output format of the exposed resources, defaults to '1'",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"resource": schema.StringAttribute{
														Description:         "Required: resource to select",
														MarkdownDescription: "Required: resource to select",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"secret_key_ref": schema.SingleNestedAttribute{
												Description:         "Selects a key of a secret in the pod's namespace",
												MarkdownDescription: "Selects a key of a secret in the pod's namespace",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key of the secret to select from. Must be a valid secret key.",
														MarkdownDescription: "The key of the secret to select from. Must be a valid secret key.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
														MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"optional": schema.BoolAttribute{
														Description:         "Specify whether the Secret or its key must be defined",
														MarkdownDescription: "Specify whether the Secret or its key must be defined",
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
								Required: false,
								Optional: true,
								Computed: false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"repeats": schema.Int64Attribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"schedule": schema.StringAttribute{
						Description:         "schedule in cron job format for scheduled test execution",
						MarkdownDescription: "schedule in cron job format for scheduled test execution",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"steps": schema.ListNestedAttribute{
						Description:         "Batch steps is list of batch tests which will be sequentially orchestrated for parallel tests in each batch",
						MarkdownDescription: "Batch steps is list of batch tests which will be sequentially orchestrated for parallel tests in each batch",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"download_artifacts": schema.SingleNestedAttribute{
									Description:         "options to download artifacts from previous steps",
									MarkdownDescription: "options to download artifacts from previous steps",
									Attributes: map[string]schema.Attribute{
										"all_previous_steps": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"previous_step_numbers": schema.ListAttribute{
											Description:         "previous step numbers starting from 1",
											MarkdownDescription: "previous step numbers starting from 1",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"previous_test_names": schema.ListAttribute{
											Description:         "previous test names",
											MarkdownDescription: "previous test names",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"execute": schema.ListNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"delay": schema.StringAttribute{
												Description:         "delay duration in time units",
												MarkdownDescription: "delay duration in time units",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"execution_request": schema.SingleNestedAttribute{
												Description:         "TestSuiteStepExecutionRequest contains parameters to be used by the executions. These fields will be passed to the execution when a Test Suite is queued for execution. TestSuiteStepExecutionRequest parameters have the highest priority. They override the values coming from Test Suites, Tests, and Test Executions.",
												MarkdownDescription: "TestSuiteStepExecutionRequest contains parameters to be used by the executions. These fields will be passed to the execution when a Test Suite is queued for execution. TestSuiteStepExecutionRequest parameters have the highest priority. They override the values coming from Test Suites, Tests, and Test Executions.",
												Attributes: map[string]schema.Attribute{
													"args": schema.ListAttribute{
														Description:         "additional executor binary arguments",
														MarkdownDescription: "additional executor binary arguments",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"args_mode": schema.StringAttribute{
														Description:         "usage mode for arguments",
														MarkdownDescription: "usage mode for arguments",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("append", "override", "replace"),
														},
													},

													"command": schema.ListAttribute{
														Description:         "executor binary command",
														MarkdownDescription: "executor binary command",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"cron_job_template": schema.StringAttribute{
														Description:         "cron job template extensions",
														MarkdownDescription: "cron job template extensions",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"cron_job_template_reference": schema.StringAttribute{
														Description:         "cron job template extensions reference",
														MarkdownDescription: "cron job template extensions reference",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"disable_webhooks": schema.BoolAttribute{
														Description:         "whether webhooks should be called on execution Deprecated: field is not used",
														MarkdownDescription: "whether webhooks should be called on execution Deprecated: field is not used",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"execution_labels": schema.MapAttribute{
														Description:         "test execution labels",
														MarkdownDescription: "test execution labels",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"http_proxy": schema.StringAttribute{
														Description:         "http proxy for executor containers",
														MarkdownDescription: "http proxy for executor containers",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"https_proxy": schema.StringAttribute{
														Description:         "https proxy for executor containers",
														MarkdownDescription: "https proxy for executor containers",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"job_template": schema.StringAttribute{
														Description:         "job template extensions",
														MarkdownDescription: "job template extensions",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"job_template_reference": schema.StringAttribute{
														Description:         "job template extensions reference",
														MarkdownDescription: "job template extensions reference",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"negative_test": schema.BoolAttribute{
														Description:         "negative test will fail the execution if it is a success and it will succeed if it is a failure",
														MarkdownDescription: "negative test will fail the execution if it is a success and it will succeed if it is a failure",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"pvc_template": schema.StringAttribute{
														Description:         "pvc template extensions",
														MarkdownDescription: "pvc template extensions",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"pvc_template_reference": schema.StringAttribute{
														Description:         "pvc template extensions reference",
														MarkdownDescription: "pvc template extensions reference",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"running_context": schema.SingleNestedAttribute{
														Description:         "RunningContext for test or test suite execution",
														MarkdownDescription: "RunningContext for test or test suite execution",
														Attributes: map[string]schema.Attribute{
															"context": schema.StringAttribute{
																Description:         "Context value depending from its type",
																MarkdownDescription: "Context value depending from its type",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"type": schema.StringAttribute{
																Description:         "One of possible context types",
																MarkdownDescription: "One of possible context types",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"scraper_template": schema.StringAttribute{
														Description:         "scraper template extensions",
														MarkdownDescription: "scraper template extensions",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"scraper_template_reference": schema.StringAttribute{
														Description:         "scraper template extensions reference",
														MarkdownDescription: "scraper template extensions reference",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"sync": schema.BoolAttribute{
														Description:         "whether to start execution sync or async",
														MarkdownDescription: "whether to start execution sync or async",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"variables": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "variable name",
																MarkdownDescription: "variable name",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"type": schema.StringAttribute{
																Description:         "variable type",
																MarkdownDescription: "variable type",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"value": schema.StringAttribute{
																Description:         "variable string value",
																MarkdownDescription: "variable string value",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"value_from": schema.SingleNestedAttribute{
																Description:         "or load it from var source",
																MarkdownDescription: "or load it from var source",
																Attributes: map[string]schema.Attribute{
																	"config_map_key_ref": schema.SingleNestedAttribute{
																		Description:         "Selects a key of a ConfigMap.",
																		MarkdownDescription: "Selects a key of a ConfigMap.",
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "The key to select.",
																				MarkdownDescription: "The key to select.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"optional": schema.BoolAttribute{
																				Description:         "Specify whether the ConfigMap or its key must be defined",
																				MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},
																		},
																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"field_ref": schema.SingleNestedAttribute{
																		Description:         "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
																		MarkdownDescription: "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
																		Attributes: map[string]schema.Attribute{
																			"api_version": schema.StringAttribute{
																				Description:         "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
																				MarkdownDescription: "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"field_path": schema.StringAttribute{
																				Description:         "Path of the field to select in the specified API version.",
																				MarkdownDescription: "Path of the field to select in the specified API version.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},
																		},
																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"resource_field_ref": schema.SingleNestedAttribute{
																		Description:         "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
																		MarkdownDescription: "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
																		Attributes: map[string]schema.Attribute{
																			"container_name": schema.StringAttribute{
																				Description:         "Container name: required for volumes, optional for env vars",
																				MarkdownDescription: "Container name: required for volumes, optional for env vars",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"divisor": schema.StringAttribute{
																				Description:         "Specifies the output format of the exposed resources, defaults to '1'",
																				MarkdownDescription: "Specifies the output format of the exposed resources, defaults to '1'",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"resource": schema.StringAttribute{
																				Description:         "Required: resource to select",
																				MarkdownDescription: "Required: resource to select",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},
																		},
																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"secret_key_ref": schema.SingleNestedAttribute{
																		Description:         "Selects a key of a secret in the pod's namespace",
																		MarkdownDescription: "Selects a key of a secret in the pod's namespace",
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "The key of the secret to select from. Must be a valid secret key.",
																				MarkdownDescription: "The key of the secret to select from. Must be a valid secret key.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"optional": schema.BoolAttribute{
																				Description:         "Specify whether the Secret or its key must be defined",
																				MarkdownDescription: "Specify whether the Secret or its key must be defined",
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
														Required: false,
														Optional: true,
														Computed: false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"test": schema.StringAttribute{
												Description:         "object name",
												MarkdownDescription: "object name",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"stop_on_failure": schema.BoolAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
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
		},
	}
}

func (r *TestsTestkubeIoTestSuiteV3Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_tests_testkube_io_test_suite_v3_manifest")

	var model TestsTestkubeIoTestSuiteV3ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("tests.testkube.io/v3")
	model.Kind = pointer.String("TestSuite")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
